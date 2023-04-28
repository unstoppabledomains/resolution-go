from ctypes import *
import codecs;

lib = cdll.LoadLibrary("../shared/libawesome.dylib")

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

lib.Addr.argtypes = [GoString, GoString]
lib.Addr.restype = GoString

def addr_of(domain, ticker):
    name = GoString(bytes(domain,'UTF-8'), len(domain))
    tic = GoString(bytes(ticker,'UTF-8'), len(ticker))
    address = lib.Addr(name, tic)
    return codecs.decode(address.p, 'UTF-8')

print("address :" + addr_of("sandy.nft", "ETH"))