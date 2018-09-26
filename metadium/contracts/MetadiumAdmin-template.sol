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

    int public memberCount;
    mapping (address => Member) members;
    address memberHead;
    address memberTail;

    int public nodeCount;
    mapping(bytes => Node) nodes;
    bytes nodeHead;
    bytes nodeTail;

    constructor() public {
        tokens = 1000000;	// To Be Substituted
        blocksPer = 10;		// To Be Substituted

        address[1] memory _members = [ msg.sender ]; // To Be Substituted
        int[1] memory _stakes = [ int(1000000) ]; // To Be Substituted
        Node[] memory _nodes; // To Be Substituted

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
        modifiedBlock = block.number;
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

    function append(BytesBuffer bb, bytes b) internal pure {
        if (bb.ix + b.length <= bb.buffer.length) {
            for (uint i = 0; i < b.length; i++)
                bb.buffer[bb.ix + i] = b[i];
        }
        bb.ix += b.length;
    }

    function bytesNull(bytes b) internal pure returns (bool) {
        return b.length == 0;
    }

    function bytesEqual(bytes a, bytes b) internal pure returns (bool) {
        return keccak256(a) == keccak256(b);
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

    // in web3 environment, convert bytes to string using web3.toAscii().
    function getNode(bytes _id) public view returns (bool present, string json) {
        Node memory n = nodes[_id];
        if (_id.length == 0 || !bytesEqual(n.id, _id))
            return;
        BytesBuffer memory bb = BytesBuffer(0, new bytes(1024));
        present = true;
        append(bb, '{"partner":');
        append(bb, bytes(n.partner ? 'true' : 'false'));
        append(bb, ',"name":"');
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
}

/* EOF */
