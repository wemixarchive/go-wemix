#!/bin/bash

[ "${SOL_GAS}" = "" ] && SOL_GAS="0x10000000"

SOLC=
solc --version > /dev/null 2>&1
if [ $? = 0 ]; then
    SOLC=solc
else
#    docker --version > /dev/null 2>&1 && SOLC="docker run -v $(pwd):/tmp --workdir /tmp --rm ethereum/solc:stable"
    docker --version > /dev/null 2>&1 && SOLC="docker run -v $(pwd):/tmp --workdir /tmp --rm ethereum/solc:0.4.24"
fi

if [ "$SOLC" = "" ]; then
    echo "Cannot find solc or docker."
    exit 1
fi

function usage ()
{
    echo "$(basename $0) [-f <format>] [-g gas] [-p gas-price] [-l <name>:<addr>]+
	<sol-file> [<js-file> | <json-file>]

-f <format>:      oupput format: \"js\" || \"json\", default is \"js\".
-g <gas>:         gas amount to spend.
-p <gas-price>:   gas price
-l <name>:<addr>: library name and address pair separated by ':'.
    Multiple -l options can be used to specify multiple libraries.
-r <path>=<path>: remap paths, e.g. \"wherever=/foo/bar\"

Output Formats:
  \"js\":   creates 'remix'-generated .js style file that can be loaded to
    geth/gmet console.
  \"json\": creates 'truffle'-generated .json style file.

Environment Variables:
  SOL_GAS for gas amount, equivalent to -g option.
  SOL_LIBS for libaries, equivalent to -l option. Pairs should be separated by
    space, e.g. \"name1:0x123..456 name2:abc..def\".
"
}

# int compile(string solFile, string jsFile)
function compile ()
{
    ${SOLC} ${PATH_REMAP} --optimize --abi --bin $1 | awk -v gas="$SOL_GAS" -v gas_price="$SOL_GASPRICE" -v libs="$SOL_LIBS" -v outfmt=$outfmt '
function flush2js() {
  if (length(gas_price) != 0) {
      gas_price_2 = ",\
    gasPrice: \"" gas_price "\"";
  }
  if (length(code_name) > 0) {
    printf "\n\
function %s_new() {\n\
  return %s_contract.new(\n\
  {\n\
    from: web3.eth.accounts[0],\n\
    data: %s_data,\n\
    gas: \"%s\"%s\n\
  }, function (e, contract) {\n\
    if (typeof contract.address !== \"undefined\") {\n\
      console.log(\"Contract mined! address: \" + contract.address + \" transactionHash: \" + contract.transactionHash);\n\
    }\n\
  });\n\
}\n\
\n\
function %s_load(addr) {\n\
   return %s_contract.at(addr);\n\
}\n\
", code_name, code_name, code_name, gas, gas_price_2, code_name, code_name;
  }
}

function flush2json() {
  if (length(code_name) > 0) {
    printf "{\n\
  \"contractName\": \"%s\",\n\
  \"abi\": %s,\n\
  \"bytecode\": \"0x%s\"\n\
}\n\
", code_name, abi, code;
  }
}

function flush2abi() {
  if (length(code_name) > 0) {
    printf "{\n\
  \"contractName\": \"%s\",\n\
  \"abi\": %s\n\
}\n\
", code_name, abi, code;
  }
}

function flush() {
  if (outfmt == "js")
    flush2js();
  else if (outfmt == "json")
    flush2json();
  else
    flush2abi();
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
    abi = $0
    if (outfmt == "js")
      print "var " code_name "_contract = web3.eth.contract(" abi ");";
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
    if (outfmt == "js")
      print "var " code_name "_data = \"0x" code "\";";
  }
}
' > $2;
}

args=`getopt f:g:l:p:r: $*`
if [ $? != 0 ]; then
    usage;
    exit 1;
fi
set -- $args

outfmt=js
for i; do
    case "$i" in
    -f)
	outfmt=$2
	shift;
	shift;;
    -g)
	SOL_GAS=$2
	shift;
	shift;;
    -p)
	SOL_GASPRICE=$2
	shift;
	shift;;
    -l)
	[ "$SOL_LIBS" = "" ] || SOL_LIBS="$SOL_LIBS "
	SOL_LIBS="${SOL_LIBS}$2";
	shift;
	shift;;
    -r)
        [ "$PATH_REMAP" = "" ] || PATH_REMAP="$PATH_REMAP "
	PATH_REMAP="${PATH_REMAP}$2";
	shift;
	shift;;
    esac
done

if [ $# != 3 -o "$outfmt" != "js" -a "$outfmt" != "json" -a "$outfmt" != "abi" ]; then
    usage
    exit 1
fi

compile "$2" "$3"

# EOF
