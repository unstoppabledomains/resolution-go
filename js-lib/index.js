const ffi = require("ffi-napi");
const ref = require("ref-napi");
const Struct = require("ref-struct-napi");

const GoString = Struct({
  p: "string",
  n: "longlong",
});

var libm = ffi.Library("../shared/libawesome", {
  Addr: [GoString, [GoString, GoString]],
});

const addrOf = async (domain, ticker) => {
  const domainString = new GoString();
  domainString["p"] = domain;
  domainString["n"] = domain.length;

  const tickerString = new GoString();
  tickerString["p"] = ticker;
  tickerString["n"] = ticker.length;

  const address = await libm.Addr(domainString, tickerString);
  return address.p;
};

addrOf("tu-nguyen.crypto", "ETH").then((data) => {
  console.log(data);
});

module.exports = {
  addrOf,
};
