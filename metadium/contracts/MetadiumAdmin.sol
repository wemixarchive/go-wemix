/* MetadiumAdmin.sol */

pragma solidity ^0.4.24;

contract Admin {
    int public magic = 0x4d6574616469756d;

    int tokens = 1000000;

    // to pre-determine miner: height / blocksPer % node-count
    int public blocksPer = 10;

    // block height when nodes update happened most recently
    uint public modifiedBlock;

    struct BytesBuffer {
        uint ix;
        bytes buffer;
    }

    struct Member {
        address addr;
        int tokens;
        address prev;
        address next;
    }

    struct TokenDistribution {
        address addr;
        int tokens;
    }

    struct MemberBallot {
        bool join_or_dismissal; // true: join, false: dismissal
        address proposer;
        address nominee;
        uint due;

        TokenDistribution []tokens;

        address []accepts;
        address []rejects;

        address prev;
        address next;
    }

    struct Node {
        bool partner;           // true: partner, false: associated/read-only
        bytes name;
        bytes id;               // admin.nodeInfo.id is 512 bit public key
        bytes ip;
        uint port;
        uint notBefore;         // block height, not used yet
        uint notAfter;

        bytes prev;
        bytes next;
    }

    struct NodeBallot {
        bool join_or_dismissal;
        address proposer;
        bytes name;
        bytes id;
        bytes ip;
        uint port;
        uint notBefore;         // block height, not used yet
        uint notAfter;
        uint due;

        address[] accepts;
        address[] rejects;

        bytes prev;
        bytes next;
    }

    int public memberCount;
    mapping (address => Member) members;
    address memberHead;
    address memberTail;

    int public memberBallotCount;
    mapping (address => MemberBallot) memberBallots;
    address memberBallotHead;
    address memberBallotTail;

    int public nodeCount;
    mapping(bytes => Node) nodes;
    bytes nodeHead;
    bytes nodeTail;

    int public nodeBallotCount;
    mapping (bytes => NodeBallot) nodeBallots;
    bytes nodeBallotHead;
    bytes nodeBallotTail;

    event MemberJoined(address addr);
    event MemberDismissed(address addr);
    event NodeJoined(bytes name, bytes id, bytes ip, uint port);
    event NodeDismissed(bytes name, bytes id, bytes ip, uint port);

    constructor() public {
        tokens = 1000000;	// To Be Substituted
        blocksPer = 10;		// To Be Substituted

        address[1] memory _members = [ msg.sender ]; // To Be Substituted
        int[1] memory _stakes = [ int(1000000) ]; // To Be Substituted
        Node[1] memory _nodes = [ Node(true, "dummy1", "53806c29df593f8699de59b2205db61cba2e61473f29710d63469e1eb04ac35bb60ef796e70a29248ebdcbaabaea1f902fbdea1b514846c8d4374b93d387e349", "172.18.0.1", 10009, 0, 0, "", "") ]; // To Be Substituted

        uint i;
        for (i = 0; i < _members.length; i++) {
            address a = _members[i];
            members[a].addr = a;
            members[a].tokens = _stakes[i];
            memberLinkAppend(false, a);
            memberCount++;
        }

        for (i = 0; i < _nodes.length; i++) {
            Node memory n = _nodes[i];
            nodes[n.id] = n;
            nodeLinkAppend(false, n.id);
            nodeCount++;
        }
    }

    function atoi(bytes bs) internal pure returns (bool, int) {
        int v = 0;
        for (uint i = 0; i < bs.length; i++) {
            if (bs[i] < 0x30 || bs[i] > 0x39)   // '0' = 0x30 , '9' = 0x39
                return (false, 0);
            v = v * 10 + (int(bs[i]) - 0x30);
        }
        return (true, v);
    }

    function htoi(bytes bs) public pure returns (bool, uint) {
        uint v = 0;
        uint i = 0;
        if (bs.length >= 2 && bs[0] == '0' && (bs[1] == 'x' || bs[1] == 'X'))
            i = 2;
        for ( ; i < bs.length; i++) {
            uint j = uint(bs[i]);
            uint k;

            // 0=0x30, 9=0x39, A=0x41, F=0x46, a=0x61, f=0x61
            if (j >= 0x30 && j <= 0x39)
                k = j - 0x30;
            else if (j >= 0x41 && j <= 0x46)
                k = 10 + j - 0x41;
            else if (j >= 0x61 && j <= 0x66)
                k = 10 + j - 0x61;
            else {
                return (false, i);
            }

            v = v * 16 + k;
        }
        return (true, v);
    }

    function itoa(int v) internal pure returns (bytes) {
        uint l = 1;
        int vv = v;
        if (v < 0) {
            l++;
            vv = -v;
        }
        while (vv >= 10) {
            l++;
            vv /= 10;
        }

        bytes memory s = new bytes(l);
        l--;
        vv = v;
        if (v < 0) {
            s[0] = 0x2d;                // '-'
            vv = -v;
        }
        do {
            s[l--] = bytes1(0x30 + (vv % 10));  // 0x30 == '0'
            vv /= 10;
        } while (vv > 0);
        return s;
    }

    // bytes4 -> bytes
    function b4b(bytes4 x) internal pure returns (bytes) {
        uint l = 0;
        for (uint i = 0; i < 4; i++) {
            if (x[i] == 0)
                break;
            l++;
        }
        bytes memory y = new bytes(l);
        for (i = 0; i < l; i++) {
            y[i] = x[i];
        }
        return y;
    }

    // address to bytes
    function addrtob(address addr) internal pure returns (bytes baddr) {
        assembly {
            let m := mload(0x40)
            mstore(add(m, 20), xor(0x140000000000000000000000000000000000000000, addr))
            mstore(0x40, add(m, 52))
            baddr := m
        }
    }

    function nibble2hex(uint8 n, bytes bb, uint ix) internal pure {
        if (n < 10) {
            bb[ix] = byte(0x30 + n);
        } else {
            bb[ix] = byte(0x61 + n - 10);
        }
    }

    // address to hex string bytes
    function addrtoh(address addr) internal pure returns (bytes baddr) {
        baddr = new bytes(42);
        baddr[0] = '0';
        baddr[1] = 'x';
        uint uaddr = uint(addr);
        for (uint i = 0; i < 20; i++) {
            uint8 b = uint8(uaddr / (2**(8*(19 - i))));
            nibble2hex(b / 16, baddr, 2 + i * 2);
            nibble2hex(b % 16, baddr, 2 + i * 2 + 1);
        }
    }

    function append(BytesBuffer bb, bytes b) internal pure {
        if (bb.ix + b.length <= bb.buffer.length) {
            for (uint i = 0; i < b.length; i++)
                bb.buffer[bb.ix + i] = b[i];
        }
        bb.ix += b.length;
    }

    function appendArray(BytesBuffer bb, bytes[] bs) internal pure {
        for (uint i = 0; i < bs.length; i++) {
            append(bb, bs[i]);
        }
    }

    function bytesNull(bytes b) internal pure returns (bool) {
        return b.length == 0;
    }

    function bytesEqual(bytes a, bytes b) internal pure returns (bool) {
        return keccak256(a) == keccak256(b);
    }

    function getToken(bytes bs, uint ix, uint len, bytes1 delimiter) internal pure returns (uint new_ix, bytes token) {
        uint six = ix;
        while (ix < len) {
            if (bs[ix++] == delimiter)
                break;
        }
        new_ix = ix;
        uint l = ix - six;
        if (ix < len - 1)
            l--;
        token = new bytes(l);
        for (uint i = 0; i < l; i++) {
            token[i] = bs[six + i];
        }
    }

    function memberLinkAppend(bool is_ballot, address addr) internal {
        if (!is_ballot) {
            if (memberHead == 0)
                memberHead = addr;
            if (memberTail == 0)
                memberTail = addr;
            else {
                members[memberTail].next = addr;
                members[addr].prev = memberTail;
                memberTail = addr;
            }
        }
        else {
            if (memberBallotHead == 0)
                memberBallotHead = addr;
            if (memberBallotTail == 0)
                memberBallotTail = addr;
            else {
                memberBallots[memberBallotTail].next = addr;
                memberBallots[addr].prev = memberBallotTail;
                memberBallotTail = addr;
            }
        }
    }

    function memberLinkDelete(bool is_ballot, address addr) internal {
        if (!is_ballot) {
            if (members[addr].next != 0)
                members[members[addr].next].prev = members[addr].prev;
            if (members[addr].prev != 0)
                members[members[addr].prev].next = members[addr].next;
            if (memberHead == addr)
                memberHead = members[addr].next;
            if (memberTail == addr)
                memberTail = members[addr].prev;
        }
        else {
            if (memberBallots[addr].next != 0)
                memberBallots[memberBallots[addr].next].prev = memberBallots[addr].prev;
            if (memberBallots[addr].prev != 0)
                memberBallots[memberBallots[addr].prev].next = memberBallots[addr].next;
            if (memberBallotHead == addr)
                memberBallotHead = memberBallots[addr].next;
            if (memberBallotTail == addr)
                memberBallotTail = memberBallots[addr].prev;
        }
    }

    function nodeLinkAppend(bool is_ballot, bytes id) internal {
        if (!is_ballot) {
            if (bytesNull(nodeHead))
                nodeHead = id;
            if (bytesNull(nodeTail))
                nodeTail = id;
            else {
                nodes[nodeTail].next = id;
                nodes[id].prev = nodeTail;
                nodeTail = id;
            }
        }
        else {
            if (bytesNull(nodeBallotHead))
                nodeBallotHead = id;
            if (bytesNull(nodeBallotTail))
                nodeBallotTail = id;
            else {
                nodeBallots[nodeBallotTail].next = id;
                nodeBallots[id].prev = nodeBallotTail;
                nodeBallotTail = id;
            }
        }
    }

    function nodeLinkDelete(bool is_ballot, bytes id) internal {
        if (!is_ballot) {
            if (!bytesNull(nodes[id].next))
                nodes[nodes[id].next].prev = nodes[id].prev;
            if (!bytesNull(nodes[id].prev))
                nodes[nodes[id].prev].next = nodes[id].next;
            if (bytesEqual(nodeHead, id))
                nodeHead = nodes[id].next;
            if (bytesEqual(nodeTail, id))
                nodeTail = nodes[id].prev;
        }
        else {
            if (!bytesNull(nodeBallots[id].next))
                nodeBallots[nodeBallots[id].next].prev = nodeBallots[id].prev;
            if (!bytesNull(nodeBallots[id].prev))
                nodeBallots[nodeBallots[id].prev].next = nodeBallots[id].next;
            if (bytesEqual(nodeBallotHead, id))
                nodeBallotHead = nodeBallots[id].next;
            if (bytesEqual(nodeBallotTail, id))
                nodeBallotTail = nodeBallots[id].prev;
        }
    }

    function getAddr(bytes bs, uint ix) internal view returns (uint new_ix, bool success, address addr) {
        bytes memory baddr;
        (new_ix, baddr) = getToken(bs, ix, bs.length, bytes1(":"));
        if (new_ix >= bs.length || baddr.length == 0)
            return;
        uint iaddr;
        (success, iaddr) = htoi(baddr);
        if (!success)
            return;
        success = false;
        addr = address(iaddr);
        if (members[addr].addr != addr)
            return;
        success = true;
    }

    function getAmount(bytes bs, uint ix) internal pure returns (uint new_ix, bool success, int amount) {
        bytes memory samt;
        (new_ix, samt) = getToken(bs, ix, bs.length, bytes1(","));
        if (samt.length == 0)
            return;
        (success, amount) = atoi(samt);
        if (amount <= 0)
            success = false;
        return;
    }

    // string tokens is delimited by '/', e.g. "<addr1>/100/<addr2>/200"
    function proposeMemberBallot(bool join_or_dismissal, address nominee, string _tokens, uint due) internal returns (bool rc, string reason) {
        if (members[msg.sender].tokens <= 0) {
            reason = "Not Allowed";
            return;
        }
        if (join_or_dismissal == true) {
            if (members[nominee].addr == nominee) {
                reason = "Already Member";
                return;
            }
        }
        else {
            if (members[nominee].addr == 0) {
                reason = "Not a Member";
                return;
            }
        }
        if (memberBallots[nominee].nominee == nominee) {
            reason = "Already Proposed";
            return;
        }
        if (due < block.number + 1) {
            reason = "Wrong Due";
            return;
        }

        TokenDistribution[] memory tds;
        bytes memory bs = bytes(_tokens);
        uint ix;
        uint i;
        bool b;
        address addr;
        int amt;

        ix = 0;
        i = 0;
        while (ix < bs.length) {
            (ix, b, addr) = getAddr(bs, ix);
            if (!b) {
                reason = "Invalid Address";
                return;
            }
            (ix, b, amt) = getAmount(bs, ix);
            if (!b) {
                reason = "Invalid Amount";
                return;
            }
            if (join_or_dismissal) {
                if (members[addr].tokens < amt) {
                    reason = "Not Enough Tokens";
                    return;
                }
            }

            i++;
        }

        tds = new TokenDistribution[](i);
        ix = 0;
        i = 0;
        while (ix < bs.length) {
            (ix, b, addr) = getAddr(bs, ix);
            (ix, b, amt) = getAmount(bs, ix);
            tds[i++] = TokenDistribution(addr, amt);
        }

        if (!join_or_dismissal) {
            amt = 0;
            for (i = 0; i < tds.length; i++) {
                amt += tds[i].tokens;
            }
            if (amt != members[nominee].tokens) {
                reason = "Amount of Token Mismatch";
                return;
            }
        }

        memberBallots[nominee].join_or_dismissal = join_or_dismissal;
        memberBallots[nominee].proposer = msg.sender;
        memberBallots[nominee].nominee = nominee;
        memberBallots[nominee].due = due;
        for (i = 0; i < tds.length; i++) {
            memberBallots[nominee].tokens.push(tds[i]);
        }

        memberLinkAppend(true, nominee);
        memberBallotCount++;

        rc = true;
    }

    function processMemberBallot(bool is_accept, address nominee) internal returns (bool rc, string reason) {
        MemberBallot memory mb = memberBallots[nominee];
        if (nominee == 0 || mb.nominee != nominee) {
            reason = "Not a Nominee";
            return;
        }

        uint i;
        uint l;
        bool is_in = false;

        if (is_accept) {
            for (i = 0; i < mb.rejects.length; i++) {
                if (mb.rejects[i] == msg.sender) {
                    if (i < mb.rejects.length-1) {
                        l = memberBallots[nominee].rejects.length - 1;
                        memberBallots[nominee].rejects[i] =
                            memberBallots[nominee].rejects[l];
                    }
                    memberBallots[nominee].rejects.length--;
                    break;
                }
            }

            for (i = 0; i < mb.accepts.length; i++) {
                if (mb.accepts[i] == msg.sender) {
                    is_in = true;
                    break;
                }
            }
        }
        else {
            for (i = 0; i < mb.accepts.length; i++) {
                if (mb.accepts[i] == msg.sender) {
                    if (i < mb.accepts.length-1) {
                        l = memberBallots[nominee].accepts.length - 1;
                        memberBallots[nominee].accepts[i] =
                            memberBallots[nominee].accepts[l];
                    }
                    memberBallots[nominee].accepts.length--;
                    break;
                }
            }

            for (i = 0; i < mb.rejects.length; i++) {
                if (mb.rejects[i] == msg.sender) {
                    is_in = true;
                    break;
                }
            }
        }

        if (!is_in) {
            if (is_accept)
                memberBallots[nominee].accepts.push(msg.sender);
            else
                memberBallots[nominee].rejects.push(msg.sender);
        }

        (rc, reason) = settleMemberBallot(nominee);
    }

    function nominateMember(address nominee, string _tokens, uint due) public returns (bool rc, string reason) {
        (rc, reason) = proposeMemberBallot(true, nominee, _tokens, due);
    }

    function acceptMember(address nominee) public returns (bool rc, string reason) {
        (rc, reason) = processMemberBallot(true, nominee);
    }

    function rejectMember(address nominee) public returns (bool rc, string reason) {
        (rc, reason) = processMemberBallot(false, nominee);
    }

    // tokens are distributed equally
    function dismissMember(address nominee, string _tokens, uint due) public returns (bool rc, string reason) {
        (rc, reason) = proposeMemberBallot(false, nominee, _tokens, due);
    }

    function acceptMemberDismissal(address nominee) public returns (bool rc, string reason) {
        (rc, reason) = processMemberBallot(true, nominee);
    }

    function rejectMemberDismissal(address nominee) public returns (bool rc, string reason) {
        (rc, reason) = processMemberBallot(false, nominee);
    }

    function getMember(address _addr) public view returns (bool present, address addr, int _tokens, address prev, address next) {
        Member memory m = members[_addr];
        if (_addr == 0 || m.addr != _addr) {
            return;
        }
        present = true;
        addr = m.addr;
        _tokens = m.tokens;
        prev = m.prev;
        next = m.next;
    }

    function firstMember() public view returns (bool present, address addr, int _tokens, address prev, address next) {
        (present, addr, _tokens, prev, next) = getMember(memberHead);
    }

    function lastMember() public view returns (bool present, address addr, int _tokens, address prev, address next) {
        (present, addr, _tokens, prev, next) = getMember(memberTail);
    }

    function getMemberBallot(address _addr) public view returns (bool present, string json) {
        MemberBallot memory m = memberBallots[_addr];
        if (m.nominee != _addr)
            return;

        present = true;

        BytesBuffer memory bb = BytesBuffer(0, new bytes(1024));
        uint ix = 0;

        bytes[] memory x = new bytes[](13);
        x[ix++] = '{"join":';
        if (m.join_or_dismissal)
            x[ix++] = 'true';
        else
            x[ix++] = 'false';
        x[ix++] = ',"proposer":"';
        x[ix++] = addrtoh(m.proposer);
        x[ix++] = '","nominee":"';
        x[ix++] = addrtoh(m.nominee);
        x[ix++] = '","due":';
        x[ix++] = itoa(int(m.due));
        x[ix++] = ',"prev":"';
        x[ix++] = addrtoh(m.prev);
        x[ix++] = '","next":"';
        x[ix++] = addrtoh(m.next);
        x[ix++] = '","tokens":[';
        appendArray(bb, x);

        uint i = 0;
        i = m.tokens.length * 4 + 1 + m.accepts.length * 3 + 1 + m.rejects.length * 3 + 1;
        x = new bytes[](i);
        ix = 0;
        for (i = 0; i < m.tokens.length; i++) {
            if (i != 0)
                x[ix++] = ',"';
            else
                x[ix++] = '"';
            x[ix++] = addrtoh(m.tokens[i].addr);
            x[ix++] = '",';
            x[ix++] = itoa(m.tokens[i].tokens);
        }

        x[ix++] = '],"accepts":[';

        for (i = 0; i < m.accepts.length; i++) {
            if (i != 0)
                x[ix++] = ',"';
            else
                x[ix++] = '"';
            x[ix++] = addrtoh(m.accepts[i]);
            x[ix++] = '"';
        }

        x[ix++] = '],"rejects":[';

        for (i = 0; i < m.rejects.length; i++) {
            if (i != 0)
                x[ix++] = ',"';
            else
                x[ix++] = '"';
            x[ix++] = addrtoh(m.rejects[i]);
            x[ix++] = '"';
        }

        x[ix++] = ']}';
        appendArray(bb, x);

        json = string(bb.buffer);
    }

    function firstMemberBallot() public view returns (bool present, string json) {
        return getMemberBallot(memberBallotHead);
    }

    function lastMemberBallot() public view returns (bool present, string json) {
        return getMemberBallot(memberBallotTail);
    }

    // validate membership and fund
    function validateMemberBallot(bool join_or_dismissal, address nominee) internal view returns (bool rc, string reason) {
        MemberBallot memory mb = memberBallots[nominee];
        Member memory m;
        uint i;
        for (i = 0; i < mb.accepts.length; i++) {
            m = members[mb.accepts[i]];
            if (m.addr == 0) {
                reason = "Accepted by Unknown Member";
                return;
            }
        }
        for (i = 0; i < mb.rejects.length; i++) {
            m = members[mb.rejects[i]];
            if (m.addr == 0) {
                reason = "Rejected by Unknown Member";
                return;
            }
        }
        TokenDistribution memory t;
        int tt;
        for (i = 0; i < mb.tokens.length; i++) {
            t = mb.tokens[i];
            if (members[t.addr].addr == 0) {
                reason = "Unknown Member in Token Distribution";
                return;
            }
            if (!join_or_dismissal) {
                tt += t.tokens;
            } else {
                if (members[t.addr].tokens < t.tokens) {
                    reason = "Not Enough Tokens Off Of Member";
                    return;
                }
            }
        }

        if (!join_or_dismissal && tt != m.tokens) {
            reason = "Token Distribution Incorrect for Dismissal";
            return;
        }

        rc = true;
    }

    /* need to 1) accomodate, and 2) validate  membership and fund change. */
    function settleMemberBallot(address nominee) public returns (bool rc, string reason) {
        if (members[msg.sender].addr != msg.sender || memberBallots[nominee].nominee != nominee) {
            reason = "Not Allowed";
            return;
        }

        MemberBallot memory mb = memberBallots[nominee];
        if (!mb.join_or_dismissal && members[nominee].addr != nominee) {
            // dismiss it

            memberBallotCount--;
            memberLinkDelete(true, nominee);
            delete memberBallots[nominee];
            reason = "Dismissed as it's an invalid Ballot.";
            return;
        }

        // check membership and funds
        (rc, reason) = validateMemberBallot(mb.join_or_dismissal, nominee);
        if (rc != true)
            return;

        int acceptTokens;
        int rejectTokens;
        int targetTokens = tokens / 2;
        if (targetTokens % 2 == 0)
            targetTokens += 1;

        uint i;
        for (i = 0; i < mb.accepts.length; i++) {
            acceptTokens += members[mb.accepts[i]].tokens;
        }
        for (i = 0; i < mb.rejects.length; i++) {
            rejectTokens += members[mb.rejects[i]].tokens;
        }

        BytesBuffer memory bb;
        TokenDistribution memory t;
        if (acceptTokens >= targetTokens) {
            if (mb.join_or_dismissal) {                 // join
                members[nominee].addr = nominee;

                for (i = 0; i < mb.tokens.length; i++) {
                    t = mb.tokens[i];
                    if (members[t.addr].tokens < t.tokens)
                        continue;
                    members[t.addr].tokens -= t.tokens;
                    members[nominee].tokens += t.tokens;
                }

                memberBallotCount--;
                memberCount++;
                memberLinkDelete(true, nominee);
                memberLinkAppend(false, nominee);
                delete memberBallots[nominee];

                bb = BytesBuffer(0, new bytes(256));
                append(bb, addrtoh(nominee));
                append(bb, " is promoted.");
                reason = string(bb.buffer);

                emit MemberJoined(nominee);

                rc = true;
                return;
            }
            else {              // dismissal
                for (i = 0; i < mb.tokens.length; i++) {
                    t = mb.tokens[i];
                    members[t.addr].tokens += t.tokens;
                    members[nominee].tokens -= t.tokens;
                }

                memberBallotCount--;
                memberCount--;
                memberLinkDelete(true, nominee);
                memberLinkDelete(false, nominee);
                delete memberBallots[nominee];
                delete members[nominee];

                bb = BytesBuffer(0, new bytes(256));
                append(bb, addrtoh(nominee));
                append(bb, " is dismissed.");
                reason = string(bb.buffer);

                emit MemberDismissed(nominee);

                rc = true;
                return;
            }
        }
        else if (rejectTokens >= targetTokens) {
            memberBallotCount--;
            memberLinkDelete(true, nominee);
            delete memberBallots[nominee];

            reason = "Ballot is rejected.";
            return;
        }
        else if (block.number >= mb.due) {
            // over due
            memberBallotCount--;
            memberLinkDelete(true, nominee);
            delete memberBallots[nominee];

            reason = "Ballot is dismissed as it's over due.";
            return;
        }
    }

    // nodes

    function proposeNodeBallot(bool join_or_dismissal, bytes name, bytes id, bytes ip, uint port, uint notBefore, uint notAfter, uint due) internal returns (bool rc, string reason) {
        if (members[msg.sender].tokens <= 0) {
            reason = "Not Allowed";
            return;
        }
        Node memory n;
        if (join_or_dismissal == true) {
            if (bytesEqual(nodes[id].id, id)) {
                reason = "Already Joined";
                return;
            }
            if ((notBefore != 0 && notAfter != 0 && notBefore >= notAfter) ||
                (notAfter != 0 && notAfter <= block.number + 1)) {
                reason = "Wrong NotBefore or NotAfter";
                return;
            }
        }
        else {
            if (bytesNull(nodes[id].id)) {
                reason = "Not Joined";
                return;
            }
            n = nodes[id];
            name = n.name;
            ip = n.ip;
            port = n.port;
        }
        if (bytesEqual(nodeBallots[id].id, id)) {
            reason = "Already Proposed";
            return;
        }
        if (due < block.number + 1) {
            reason = "Wrong Due";
            return;
        }

        nodeBallots[id].join_or_dismissal = join_or_dismissal;
        nodeBallots[id].proposer = msg.sender;
        nodeBallots[id].name = name;
        nodeBallots[id].id = id;
        nodeBallots[id].ip = ip;
        nodeBallots[id].port = port;
        nodeBallots[id].notBefore = notBefore;
        nodeBallots[id].notAfter = notAfter;
        nodeBallots[id].due = due;
        nodeLinkAppend(true, id);
        nodeBallotCount++;
        rc = true;
    }

    function processNodeBallot(bool is_accept, bytes id) internal returns (bool rc, string reason) {
        NodeBallot memory nb = nodeBallots[id];
        if (bytesNull(id) || !bytesEqual(nb.id, id)) {
            reason = "Not a Nominee";
            return;
        }

        uint i;
        uint l;
        bool is_in = false;

        if (is_accept) {
            for (i = 0; i < nb.rejects.length; i++) {
                if (nb.rejects[i] == msg.sender) {
                    if (i < nb.rejects.length-1) {
                        l = nodeBallots[id].rejects.length - 1;
                        nodeBallots[id].rejects[i] =
                            nodeBallots[id].rejects[l];
                    }
                    nodeBallots[id].rejects.length--;
                    break;
                }
            }

            for (i = 0; i < nb.accepts.length; i++) {
                if (nb.accepts[i] == msg.sender) {
                    is_in = true;
                    break;
                }
            }
        }
        else {
            for (i = 0; i < nb.accepts.length; i++) {
                if (nb.accepts[i] == msg.sender) {
                    if (i < nb.accepts.length-1) {
                        l = nodeBallots[id].accepts.length - 1;
                        nodeBallots[id].accepts[i] =
                            nodeBallots[id].accepts[l];
                    }
                    nodeBallots[id].accepts.length--;
                    break;
                }
            }

            for (i = 0; i < nb.rejects.length; i++) {
                if (nb.rejects[i] == msg.sender) {
                    is_in = true;
                    break;
                }
            }
        }

        if (!is_in) {
            if (is_accept)
                nodeBallots[id].accepts.push(msg.sender);
            else
                nodeBallots[id].rejects.push(msg.sender);
        }

        (rc, reason) = settleNodeBallot(id);
        return;
    }

    function nominateNode(bytes name, bytes id, bytes ip, uint port, uint notBefore, uint notAfter, uint due) public returns (bool rc, string reason) {
        (rc, reason) = proposeNodeBallot(true, name, id, ip, port, notBefore, notAfter, due);
    }

    function acceptNode(bytes id) public returns (bool rc, string reason) {
        (rc, reason) = processNodeBallot(true, id);
    }

    function rejectNode(bytes id) public returns (bool rc, string reason) {
        (rc, reason) = processNodeBallot(false, id);
    }

    function dismissNode(bytes id, uint due) public returns (bool rc, string reason) {
        (rc, reason) = proposeNodeBallot(false, "", id, "", 0, 0, 0, due);
    }

    function acceptNodeDismissal(bytes id) public returns (bool rc, string reason) {
        (rc, reason) = processNodeBallot(true, id);
    }

    function rejectNodeDismissal(bytes id) public returns (bool rc, string reason) {
        (rc, reason) = processNodeBallot(false, id);
    }

    // in web3 environment, convert bytes to string using web3.toAscii().
    function getNode(bytes _id) public view returns (bool present, string json) {
        Node memory n = nodes[_id];
        if (_id.length == 0 || !bytesEqual(n.id, _id))
            return;
        BytesBuffer memory bb = BytesBuffer(0, new bytes(1024));
        present = true;
        append(bb, '{"name":"');
        append(bb, n.name);
        append(bb, '","id":"');
        append(bb, n.id);
        append(bb, '","ip":"');
        append(bb, n.ip);
        append(bb, '","port":');
        append(bb, itoa(int(n.port)));
        append(bb, ',"notBefore":');
        append(bb, itoa(int(n.notBefore)));
        append(bb, ',"notAfter":');
        append(bb, itoa(int(n.notAfter)));
        append(bb, ',"prev":"');
        append(bb, n.prev);
        append(bb, '","next":"');
        append(bb, n.next);
        append(bb, '"}');
        json = string(bb.buffer);
    }

    function firstNode() public view returns (bool present, string json) {
        (present, json) = getNode(nodeHead);
    }

    function lastNode() public view returns (bool present, string json) {
        (present, json) = getNode(nodeTail);
    }

    function getNodeBallot(bytes id) public view returns (bool present, string json) {
        NodeBallot memory n = nodeBallots[id];
        if (!bytesEqual(n.id, id))
            return;

        present = true;

        BytesBuffer memory bb = BytesBuffer(0, new bytes(1024));
        uint ix = 0;

        bytes[] memory x = new bytes[](19);
        x[ix++] = '{"join":';
        if (n.join_or_dismissal)
            x[ix++] = 'true';
        else
            x[ix++] = 'false';
        x[ix++] = ',"proposer":"';
        x[ix++] = addrtoh(n.proposer);
        x[ix++] = '","name":"';
        x[ix++] = n.name;
        x[ix++] = '","id":"';
        x[ix++] = n.id;
        x[ix++] = '","ip":"';
        x[ix++] = n.ip;
        x[ix++] = '","port":"';
        x[ix++] = itoa(int(n.port));
        x[ix++] = '","due":';
        x[ix++] = itoa(int(n.due));
        x[ix++] = ',"prev":"';
        x[ix++] = n.prev;
        x[ix++] = '","next":"';
        x[ix++] = n.next;
        x[ix++] = '","accepts":[';
        appendArray(bb, x);

        uint i = 0;
        i = n.accepts.length * 3 + 1 + n.rejects.length * 3 + 1;
        x = new bytes[](i);
        ix = 0;

        for (i = 0; i < n.accepts.length; i++) {
            if (i != 0)
                x[ix++] = ',"';
            else
                x[ix++] = '"';
            x[ix++] = addrtoh(n.accepts[i]);
            x[ix++] = '"';
        }

        x[ix++] = '],"rejects":[';

        for (i = 0; i < n.rejects.length; i++) {
            if (i != 0)
                x[ix++] = ',"';
            else
                x[ix++] = '"';
            x[ix++] = addrtoh(n.rejects[i]);
            x[ix++] = '"';
        }

        x[ix++] = ']}';
        appendArray(bb, x);

        json = string(bb.buffer);
    }

    function firstNodeBallot() public view returns (bool present, string json) {
        return getNodeBallot(nodeBallotHead);
    }

    function lastNodeBallot() public view returns (bool present, string json) {
        return getNodeBallot(nodeBallotTail);
    }

    // validate membership
    function validateNodeBallot(bytes id) internal view returns (bool rc, string reason) {
        NodeBallot memory nb = nodeBallots[id];
        Member memory m;
        uint i;
        for (i = 0; i < nb.accepts.length; i++) {
            m = members[nb.accepts[i]];
            if (m.addr == 0) {
                reason = "Accepted by Unknown Member";
                return;
            }
        }
        for (i = 0; i < nb.rejects.length; i++) {
            m = members[nb.rejects[i]];
            if (m.addr == 0) {
                reason = "Rejected by Unknown Member";
                return;
            }
        }
        rc = true;
    }

    function settleNodeBallot(bytes id) public returns (bool rc, string reason) {
        if (members[msg.sender].addr != msg.sender || !bytesEqual(nodeBallots[id].id, id)) {
            reason = "Not Allowed";
            return;
        }

        NodeBallot memory nb = nodeBallots[id];
        if (!nb.join_or_dismissal && !bytesEqual(nodes[id].id, id)) {
            // dismiss it

            nodeBallotCount--;
            nodeLinkDelete(true, id);
            delete nodeBallots[id];
            reason = "Dismissed as it's an invalid Ballot.";
            return;
        }

        // check membership
        (rc, reason) = validateNodeBallot(id);
        if (rc != true)
            return;

        int acceptTokens;
        int rejectTokens;
        int targetTokens = tokens / 2;
        if (targetTokens % 2 == 0)
            targetTokens += 1;

        uint i;
        for (i = 0; i < nb.accepts.length; i++) {
            acceptTokens += members[nb.accepts[i]].tokens;
        }
        for (i = 0; i < nb.rejects.length; i++) {
            rejectTokens += members[nb.rejects[i]].tokens;
        }

        BytesBuffer memory bb;
        if (acceptTokens >= targetTokens) {
            if (nb.join_or_dismissal) {                 // join
                nodes[id].partner = true;
                nodes[id].name = nb.name;
                nodes[id].id = nb.id;
                nodes[id].ip = nb.ip;
                nodes[id].port = nb.port;
                nodes[id].notBefore = nb.notBefore;
                nodes[id].notAfter = nb.notAfter;

                nodeBallotCount--;
                nodeCount++;
                nodeLinkDelete(true, id);
                nodeLinkAppend(false, id);
                delete nodeBallots[id];

                bb = BytesBuffer(0, new bytes(256));
                append(bb, "Node ");
                append(bb, id);
                append(bb, " has joined.");
                reason = string(bb.buffer);

                emit NodeJoined(nb.name, nb.id, nb.ip, nb.port);
                modifiedBlock = block.number;

                rc = true;
                return;
            }
            else {              // dismissal
                nodeBallotCount--;
                nodeCount--;
                nodeLinkDelete(true, id);
                nodeLinkDelete(false, id);
                delete nodeBallots[id];
                delete nodes[id];

                bb = BytesBuffer(0, new bytes(256));
                append(bb, "Node ");
                append(bb, id);
                append(bb, " has been dismissed.");
                reason = string(bb.buffer);

                emit NodeDismissed(nb.name, nb.id, nb.ip, nb.port);
                modifiedBlock = block.number;

                rc = true;
                return;
            }
        }
        else if (rejectTokens >= targetTokens) {
            nodeBallotCount--;
            nodeLinkDelete(true, id);
            delete nodeBallots[id];

            reason = "Ballot is rejected.";
            return;
        }
        else if (block.number >= nb.due) {
            // over due
            nodeBallotCount--;
            nodeLinkDelete(true, id);
            delete nodeBallots[id];

            reason = "Ballot is dismissed as it's over due.";
            return;
        }
    }
}

/* EOF */
