#!/bin/bash

[ "${SOL_GAS}" = "" ] && SOL_GAS="0x10000000"


function usage ()
{
    echo "$(basename $0) [-g gas] [-l <name>:<addr>]+ <sol-file> <js-file>

-g <gas>:         gas amount to spend.
-l <name>:<addr>: library name and address pair separated by ':'.
    Multiple -l options can be used to specify multiple libraries.

Environment Variables:
  SOL_GAS for gas amount, equivalent to -g option.
  SOL_LIBS for libaries, equivalent to -l option. Pairs should be separated by
    space, e.g. \"name1:0x123..456 name2:abc..def\".
"
}

# int compile(string solFile, string jsFile)
function compile ()
{
    docker run -v $(pwd):/tmp --workdir /tmp --rm ethereum/solc:stable --optimize --abi --bin $1 | awk -v gas="$SOL_GAS" -v libs="$SOL_LIBS" '
function flush() {
  if (length(code_name) > 0) {
    printf "\
function %s_new() {\
  %s = %s_contract.new(\
  {\
    from: web3.eth.accounts[0],\
    data: %s_data,\
    gas: \"%s\"\
  }, function (e, contract) {\
    console.log(e, contract);\
    if (typeof contract.address !== \"undefined\") {\
      console.log(\"Contract mined! address: \" + contract.address + \" transactionHash: \" + contract.transactionHash);\
    }\
  });\
}\
\
function %s_load(addr) {\
   %s = %s_contract.at(addr);\
}\
\
", code_name, code_name, code_name, code_name, gas, code_name, code_name, code_name;
  }
}

END {
  flush()
}

/^$/ {
  flush()
  code_name = ""
}

/^=======/ {
  code_name = $0
  sub("^=.*:", "", code_name)
  sub(" =======$", "", code_name)
}

# abi
/^\[/ {
  if (length(code_name) > 0) {
    print "var " code_name "_contract = web3.eth.contract(" $0 ");";
  }
}

# binary: 60606040, 60806040 for contracts, 610eb861 for libraries
/^6[01]/ {
  if (length(code_name) > 0) {
    code = $0;
    n = split(libs, alibs, " +");
    for (i = 1; i <= n; i++) {
      if (split(alibs[i], nv, ":") != 2)
        continue;
      sub("^0x", "", nv[2]);
      gsub("_+[^_]*" nv[1] "_+", nv[2], code);
    }
    print "var " code_name "_data = \"0x" code "\";";
  }
}
' > $2;
}

args=`getopt g:l: $*`
if [ $? != 0 ]; then
    usage;
    exit 1;
fi
set -- $args

#SOL_LIBS
for i; do
    case "$i" in
    -g)
	SOL_GAS=$2
	shift;
	shift;;
    -l)
	[ "$SOL_LIBS" = "" ] || SOL_LIBS="$SOL_LIBS "
	SOL_LIBS="${SOL_LIBS}$2";
	shift;
	shift;;
    esac
done

if [ $# != 3 ]; then
    usage
    exit 1
fi

compile "$2" "$3"

# EOF
