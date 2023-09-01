# docs.mitmproxy.org/stable/addons-examples#http-modify-query-string
from mitmproxy import http

# Error
# Please open my apps to establish a connection with the server.

def request(f: http.HTTPFlow) -> None:
   # github.com/mitmproxy/mitmproxy/blob/main/mitmproxy/addons/blocklist.py
   if f.request.path.startswith('/fdfe/apps/contentSync'):
      f.kill()
   if f.request.path.startswith('/fdfe/moduleDelivery'):
      f.kill()
   if f.request.path.startswith('/fdfe/uploadDynamicConfig'):
      f.kill()
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
   if f.request.path.startswith('/play/log'):
      f.kill()
   # github.com/mitmproxy/mitmproxy/blob/main/mitmproxy/addons/modifyheaders.py
   if f.request.path.startswith('/fdfe/acquire'):
      f.request.headers.pop('Accept-Encoding', None)
      f.request.headers.pop('Accept-Language', None)
      f.request.headers.pop('Connection', None)
      f.request.headers.pop('User-Agent', None)
      f.request.headers.pop('X-DFE-Client-Id', None)
      f.request.headers.pop('X-DFE-Content-Filters', None)
      f.request.headers.pop('X-DFE-Cookie', None)
      f.request.headers.pop('X-DFE-Device-Config-Token', None)
      f.request.headers.pop('X-DFE-Encoded-Targets', None)
      f.request.headers.pop('X-DFE-MCCMNC', None)
      f.request.headers.pop('X-DFE-Network-Type', None)
      f.request.headers.pop('X-DFE-Request-Params', None)
      f.request.headers.pop('X-PS-RH', None)
      f.request.headers.pop('X-Public-Android-Id', None)
   if f.request.path.startswith('/fdfe/delivery'):
      f.request.headers.pop('Accept-Encoding', None)
      f.request.headers.pop('Accept-Language', None)
      f.request.headers.pop('Connection', None)
      f.request.headers.pop('X-DFE-Client-Id', None)
      f.request.headers.pop('X-DFE-Content-Filters', None)
      f.request.headers.pop('X-DFE-Cookie', None)
      f.request.headers.pop('X-DFE-Device-Config-Token', None)
      f.request.headers.pop('X-DFE-Encoded-Targets', None)
      f.request.headers.pop('X-DFE-MCCMNC', None)
      f.request.headers.pop('X-DFE-Network-Type', None)
      f.request.headers.pop('X-DFE-Request-Params', None)
      f.request.headers.pop('X-DFE-UserLanguages', None)
      f.request.headers.pop('X-PS-RH', None)