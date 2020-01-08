package gotalkjs
const BrowserLibSizeString = "15907"
const BrowserLibSHA1Base64 = "dQ0cuFjXScsN4hNYsNuL0es/MP4="
const BrowserLibETag = "\"dQ0cuFjXScsN4hNYsNuL0es/MP4=\""
const BrowserLibString = "!function(e){\"use strict\";var t,r={},n={},o=0,s=function(e){var t,s=e.replace(/^.\\//,\"\");if(!(t=n[s])){for(var a=0;a<o;a++)\". \";var i=r[s];i&&o<100&&(r[s]=null,n[s]={exports:{}},++o,i(n[s]),--o,n[s]=n[s].exports),t=n[s]}return t};r.EventEmitter=function(e){e.exports;function t(){}e.exports=t,t.prototype.addListener=function(e,t){if(\"function\"!=typeof t)throw TypeError(\"listener must be a function\");if(!this.__events)return Object.defineProperty(this,\"__events\",{value:{},enumerable:!1,writable:!0}),this.__events[e]=[t],this;var r=this.__events[e];return void 0===r?(this.__events[e]=[t],this):(r.push(t),this)},t.prototype.on=t.prototype.addListener,t.prototype.once=function(e,t){var r=!1,n=function(){this.removeListener(e,n),r||(r=!0,t.apply(this,arguments))};return this.on(e,n)},t.prototype.removeListener=function(e,t){var r,n=this.__events?this.__events[e]:void 0;if(void 0!==n){for(;-1!==(r=n.indexOf(t));)n.splice(r,1);return 0===n.length&&delete this.__events[e],n.length}return this},t.prototype.removeAllListeners=function(e){this.__events&&(e?delete this.__events[e]:delete this.__events)},t.prototype.listeners=function(e){return e?this.__events?this.__events[e]:void 0:this.__events},t.prototype.emit=function(e){var t=this.__events?this.__events[e]:void 0;if(void 0===t)return!1;for(var r=0,n=t.length,o=Array.prototype.slice.call(arguments,1);r!==n;++r)t[r]||console.log(\"e\",e,r,o),t[r].apply(this,o);return!0},t.mixin=function(e){for(var r=e;r;){if(r.__proto__===Object.prototype)return r.__proto__=t.prototype,e;r=r.__proto__}return e}},r.buf=function(e){var t;e.exports;if(\"undefined\"!=typeof Uint8Array){s(\"./utf8\");t=function(e){return e instanceof Uint8Array?e:new Uint8Array(e instanceof ArrayBuffer?e:new ArrayBuffer(e))}}e.exports=t},r.keepalive=function(e){e.exports;var t=s(\"./netaccess\"),r=s(\"./protocol\");e.exports=function(e,n,o,s){o?o<100&&(o=100):o=500,(!s||s<o)&&(s=5e3);var a,i,u,l,d,c=0;return a={isEnabled:!1,isConnected:!1,enable:function(){a.enabled||(a.enabled=!0,c=0,a.isConnected||i())},disable:function(){a.enabled&&(clearTimeout(l),a.enabled=!1,c=0)}},i=function(){clearTimeout(l),e.open(n,(function(t){d=new Date,t?u(t):(c=0,a.isConnected=!0,e.once(\"close\",u))}))},u=function(e){if(clearTimeout(l),a.isConnected=!1,a.enabled){if(t.available&&!t.onLine&&(\"undefined\"==typeof document||!document.location||\"localhost\"===document.location.hostname||\"127.0.0.1\"===document.location.hostname||\"[::1]\"===document.location.hostname))return t.once(\"online\",u),void(c=0);c=e?e.isGotalkProtocolError?e.code===r.ErrorTimeout?0:s:c?Math.min(s,2*c):o:Math.max(0,o-(new Date-d)),l=setTimeout(i,c)}},a}},r.netaccess=function(t){t.exports;var r,n=s(\"./EventEmitter\");void 0!==e&&e.addEventListener?(r=Object.create(n.prototype,{available:{value:!0,enumerable:!0},onLine:{value:!0,enumerable:!0,writable:!0}}),\"undefined\"!=typeof navigator&&(r.onLine=navigator.onLine),e.addEventListener(\"offline\",(function(e){r.onLine=!1,r.emit(\"offline\")})),e.addEventListener(\"online\",(function(e){r.onLine=!0,r.emit(\"online\")}))):r={available:!1,onLine:!0},t.exports=r},r.protocol=function(e){var t=e.exports,r=s(\"./buf\"),n=s(\"./utf8\");t.Version=1;var o=t.MsgTypeSingleReq=114,a=t.MsgTypeStreamReq=115,i=(t.MsgTypeStreamReqPart=112,t.MsgTypeSingleRes=82,t.MsgTypeStreamRes=83,t.MsgTypeErrorRes=69,t.MsgTypeRetryRes=101),u=t.MsgTypeNotification=110,l=t.MsgTypeHeartbeat=104,d=t.MsgTypeProtocolError=102;function c(e,t,r,n){for(var o,s=t||0,a=0,i=r.toString(16),u=n-i.length;u--;)e[s++]=48;for(;!isNaN(o=i.charCodeAt(a++));)e[s++]=o}function f(e,t){var n=r(t);return c(n,0,e,t),n}function p(e,t){return parseInt(String.fromCharCode(...e),t)}t.ErrorAbnormal=0,t.ErrorUnsupported=1,t.ErrorInvalidMsg=2,t.ErrorTimeout=3,t.HeartbeatMsgMaxLoad=65535,t.binary={makeFixnum:f,versionBuf:f(t.Version,2),parseVersion:function(e){return p(e,16)},parseMsg:function(e){var t,r,s,c,f,h=0;return f=1,(t=e[0])===l?(h=p(e.subarray(f,f+4),16),f+=4):t!==u&&t!==d&&(r=e.subarray(f,f+4),f+=4),t==o||t==a||t==u?(c=p(e.subarray(f,f+3),16),f+=3,s=n.decode(e.subarray(f,f+c)),f+=c):t===i&&(h=p(e.subarray(f,f+8),16),f+=8),{t,id:r,name:s,wait:h,size:p(e.subarray(f,f+8),16)}},makeMsg:function(e,t,o,s,a){var u,l,d=t?13:9;return o&&0!==o.length&&(d+=3+(l=n.encode(o)).length),(u=r(d))[0]=e,d=1,t&&4===t.length&&(\"string\"==typeof t?(u[1]=t.charCodeAt(0),u[2]=t.charCodeAt(1),u[3]=t.charCodeAt(2),u[4]=t.charCodeAt(3)):(u[1]=t[0],u[2]=t[1],u[3]=t[2],u[4]=t[3]),d+=4),l&&(c(u,d,l.length,3),d+=3,u.set(l,d),d+=l.length),e===i&&(c(u,d,s,8),d+=8),c(u,d,a,8),u},makeHeartbeatMsg:function(e){var t=r(13),n=1;return t[0]=l,c(t,n,e,4),c(t,n+=4,Math.round((new Date).getTime()/1e3),8),n+=8,t}};function h(e,t){var r=e.toString(16);return\"00000000\".substr(0,t-r.length)+r}t.text={makeFixnum:h,versionBuf:h(t.Version,2),parseVersion:function(e){return parseInt(e.substr(0,2),16)},parseMsg:function(e){var t,r,n,s,c=0;return s=1,(t=e.charCodeAt(0))===l?(c=parseInt(e.substr(s,4),16),s+=4):t!==u&&t!==d&&(r=e.substr(s,4),s+=4),t==o||t==a||t==u?n=e.substring(s+3,e.length-8):t==i&&(c=parseInt(e.substr(s,8),16),s+=8),{t,id:r,name:n,wait:c,size:parseInt(e.substr(e.length-8),16)}},makeMsg:function(e,t,r,o,s){var a=String.fromCharCode(e);return t&&4===t.length&&(a+=t),r&&0!==r.length&&(a+=h(n.sizeOf(r),3),a+=r),e===i&&(a+=h(o,8)),a+=h(s,8)},makeHeartbeatMsg:function(e){var t=String.fromCharCode(l);return t+=h(e,4),t+=h(Math.round((new Date).getTime()/1e3),8)}}},r.utf8=function(e){var t=e.exports;function r(e){for(var t,r=0,n=0;t=e.charCodeAt(n++);r+=t>>11?3:t>>7?2:1);return r}function n(e){return 255&e}if(t.sizeOf=r,\"undefined\"!=typeof TextDecoder){var o=new TextDecoder(\"utf8\"),s=new TextEncoder(\"utf8\");t.decode=function(e){return o.decode(e)},t.encode=function(e){return s.encode(e)}}else t.decode=function(e){var t,r,o=0,s=e.length-o,a=\"\";for(o=0;o<s;)(r=n(t=e[o++]))<128||(r>>5==6?t=(t<<6&2047)+(63&e[o++]):r>>4==14?(t=(t<<12&65535)+(n(e[o++])<<6&4095),t+=63&e[o++]):r>>3==30&&(t=(t<<18&2097151)+(n(e[o++])<<12&262143),t+=n(e[o++])<<6&4095,t+=63&e[o++])),a+=String.fromCharCode(t);return a},t.encode=function(e){for(var t,n=0,o=e.length,s=0,a=new Uint8Array(r(e));n!==o;)(t=e.charCodeAt(n++))<128?a[s++]=t:t<2048?(a[s++]=t>>6|192,a[s++]=63&t|128):t<65536?(a[s++]=t>>12|224,a[s++]=t>>6&63|128,a[s++]=63&t|128):(a[s++]=t>>18|240,a[s++]=t>>12&63|128,a[s++]=t>>6&63|128,a[s++]=63&t|128);return a}},function(t){var r=t.exports,n=s(\"./protocol\"),o=n.text,a=n.binary,i=s(\"./buf\"),u=s(\"./utf8\"),l=s(\"./EventEmitter\"),d=s(\"./keepalive\"),c=r;function f(e){try{return JSON.parse(e)}catch(e){}}function p(e){return Object.create(p.prototype,{handlers:{value:e,enumerable:!0},protocol:{value:i?n.binary:n.text,enumerable:!0,writable:!0},heartbeatInterval:{value:2e4,enumerable:!0,writable:!0},ws:{value:null,writable:!0},keepalive:{value:null,writable:!0},nextOpID:{value:0,writable:!0},nextStreamID:{value:0,writable:!0},pendingRes:{value:{},writable:!0},hasPendingRes:{get:function(){for(var e in this.pendingRes)return!0}},pendingClose:{value:!1,writable:!0}})}c.protocol=n,c.Buf=i,p.prototype=l.mixin(p.prototype),r.Sock=p;var h={1e3:\"normal\",1001:\"going away\",1002:\"protocol error\",1003:\"unsupported\",1005:\"no status\",1006:\"abnormal\",1007:\"inconsistent\",1008:\"invalid message\",1009:\"too large\"};p.prototype.adoptWebSocket=function(e){var t=this;if(e.readyState!==WebSocket.OPEN)throw new Error(\"web socket readyState != OPEN\");e.binaryType=\"arraybuffer\",t.ws=e,e.onclose=function(r){var n=e._gotalkCloseError;n||1e3===r.code||(n=new Error(\"websocket closed: \"+(h[r.code]||\"#\"+r.code))),function(e,t){if(e.pendingClose=!1,e.stopSendingHeartbeats(),e.ws&&(e.ws.onmessage=null,e.ws.onerror=null,e.ws.onclose=null,e.ws=null),e.nextOpID=0,e.hasPendingRes){var r=t||new Error(\"connection closed\");for(var n in e.pendingRes)e.pendingRes[n](r);e.pendingRes={}}}(t,n),t.emit(\"close\",n)},e.onmessage=function(t){e._bufferedMessages||(e._bufferedMessages=[]),e._bufferedMessages.push(t.data)}},p.prototype.adopt=function(e){if(adopt instanceof WebSocket)return this.adoptWebSocket(e);throw new Error(\"unsupported transport\")},p.prototype.handshake=function(){this.ws.send(this.protocol.versionBuf)},p.prototype.end=function(){var e=this;e.keepalive&&(e.keepalive.disable(),e.keepalive=null),!e.pendingClose&&e.hasPendingRes?e.pendingClose=!0:e.ws&&e.ws.close(1e3)},p.prototype.address=function(){return this.ws?this.ws.url:null};var g=r.ErrAbnormal=Error(\"unsupported protocol\");g.isGotalkProtocolError=!0,g.code=n.ErrorAbnormal;var v=r.ErrUnsupported=Error(\"unsupported protocol\");v.isGotalkProtocolError=!0,v.code=n.ErrorUnsupported;var y=r.ErrInvalidMsg=Error(\"invalid protocol message\");y.isGotalkProtocolError=!0,y.code=n.ErrorInvalidMsg;var b=r.ErrTimeout=Error(\"timeout\");b.isGotalkProtocolError=!0,b.code=n.ErrorTimeout,p.prototype.sendHeartbeat=function(e){var t=this.protocol.makeHeartbeatMsg(Math.round(e*n.HeartbeatMsgMaxLoad));try{this.ws.send(t)}catch(e){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(e=new Error(\"socket is closed\")),e}},p.prototype.startSendingHeartbeats=function(){var e=this;if(e.heartbeatInterval<10)throw new Error(\"Sock.heartbeatInterval is too low\");clearTimeout(e._sendHeartbeatsTimer);var t=function(){clearTimeout(e._sendHeartbeatsTimer),e.sendHeartbeat(0),e._sendHeartbeatsTimer=setTimeout(t,e.heartbeatInterval)};e._sendHeartbeatsTimer=setTimeout(t,1)},p.prototype.stopSendingHeartbeats=function(){clearTimeout(this._sendHeartbeatsTimer)},p.prototype.startReading=function(){var e,t=this,r=t.ws;function s(s){if((e=\"string\"==typeof s.data?o.parseMsg(s.data):a.parseMsg(i(s.data))).t===n.MsgTypeProtocolError){var l=e.size;l===n.ErrorAbnormal?r._gotalkCloseError=g:l===n.ErrorUnsupported?r._gotalkCloseError=v:l===n.ErrorTimeout?r._gotalkCloseError=b:r._gotalkCloseError=y,r.close(4e3+l)}else 0!==e.size&&e.t!==n.MsgTypeHeartbeat?r.onmessage=u:(t.handleMsg(e),e=null)}function u(n){var o=n.data;r.onmessage=s,t.handleMsg(e,\"string\"==typeof o?o:i(o)),e=null}r.onmessage=function(e){(\"string\"==typeof e.data?o.parseVersion(e.data):a.parseVersion(i(e.data)))!==n.Version?(r._gotalkCloseError=v,t.closeError(n.ErrorUnsupported)):(r.onmessage=s,t.heartbeatInterval>0&&t.startSendingHeartbeats())},r._bufferedMessages&&(r._bufferedMessages.forEach((function(e){r.onmessage({data:e})})),r._bufferedMessages=null)};var m={};function w(e,t){var r=e.id;\"string\"!=typeof r&&(r=String.fromCharCode(...r));var o=this,s=o.pendingRes[r];e.t===n.MsgTypeStreamRes&&t&&0!==(t.length||t.size)||(delete o.pendingRes[r],o.pendingClose&&!o.hasPendingRes&&o.end()),\"function\"==typeof s&&(e.t===n.MsgTypeErrorRes?(\"string\"!=typeof t&&(t=u.decode(t)),s(new Error(t),null)):s(null,t))}p.prototype.handleMsg=function(e,t){var r=this,o=m[e.t];o?o.call(r,e,t):(r.ws&&(r.ws._gotalkCloseError=y),r.closeError(n.ErrorInvalidMsg))},m[n.MsgTypeSingleReq]=function(e,t){var r,o,s=this;if(r=s.handlers.findRequestHandler(e.name),(o=function(t){s.sendMsg(n.MsgTypeSingleRes,e.id,null,0,t)}).error=function(t){var r=t.message||String(t);s.sendMsg(n.MsgTypeErrorRes,e.id,null,0,r)},\"function\"!=typeof r)o.error('no such operation \"'+e.name+'\"');else try{r(t,o,e.name)}catch(e){\"undefined\"!=typeof console&&console.error(e.stack||e),o.error(\"internal error\")}},m[n.MsgTypeSingleRes]=w,m[n.MsgTypeStreamRes]=w,m[n.MsgTypeErrorRes]=w,m[n.MsgTypeNotification]=function(e,t){var r=this.handlers.findNotificationHandler(e.name);r&&r(t,e.name)},m[n.MsgTypeHeartbeat]=function(e){this.emit(\"heartbeat\",{time:new Date(1e3*e.size),load:e.wait})},p.prototype.sendMsg=function(e,t,r,o,s){var a=s&&\"string\"==typeof s&&this.protocol===n.binary?u.sizeOf(s):s?s.length||s.size:0,i=this,l=i.protocol.makeMsg(e,t,r,o,a);try{i.ws.send(l),0!==a&&i.ws.send(s)}catch(e){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(e=new Error(\"socket is closed\")),e}},p.prototype.closeError=function(e){var t=this;if(t.ws){try{t.ws.send(t.protocol.makeMsg(n.MsgTypeProtocolError,null,null,0,e))}catch(e){}t.ws.close(4e3+e)}};p.prototype.bufferRequest=function(e,t,r){var o=this,s=o.nextOpID++;1679616===o.nextOpID&&(o.nextOpID=0),s=s.toString(36),s=\"0000\".substr(0,4-s.length)+s,o.pendingRes[s]=r;try{o.sendMsg(n.MsgTypeSingleReq,s,e,0,t)}catch(e){delete o.pendingRes[s],r(e)}},p.prototype.bufferNotify=function(e,t){this.sendMsg(n.MsgTypeNotification,null,e,0,t)},p.prototype.request=function(e,t,r){var n;return r?n=JSON.stringify(t):r=t,this.bufferRequest(e,n,(function(e,t){var n=f(t);return r(e,n)}))},p.prototype.notify=function(e,t){var r=JSON.stringify(t);return this.bufferNotify(e,r)},\"undefined\"!=typeof Promise&&(p.prototype.requestp=function(e,t){return new Promise((r,n)=>{this.request(e,t,(e,t)=>e?n(e):r(t))})},p.prototype.bufferRequestp=function(e,t){return new Promise((r,n)=>{this.bufferRequest(e,t,(e,t)=>e?n(e):r(t))})});var E=function(e,t,r){return Object.create(E.prototype,{s:{value:e},op:{value:t,enumerable:!0},id:{value:r,enumerable:!0}})};function M(){return Object.create(M.prototype,{reqHandlers:{value:{}},reqFallbackHandler:{value:null,writable:!0},noteHandlers:{value:{}},noteFallbackHandler:{value:null,writable:!0}})}if(l.mixin(E.prototype),E.prototype.write=function(e){this.ended||(this.started?this.s.sendMsg(n.MsgTypeStreamReqPart,this.id,null,0,e):(this.started=!0,this.s.sendMsg(n.MsgTypeStreamReq,this.id,this.op,0,e)),e&&0!==e.length&&0!==e.size||(this.ended=!0))},E.prototype.end=function(){this.write(null)},p.prototype.streamRequest=function(e){var t=this,r=t.nextStreamID++;46656===t.nextStreamID&&(t.nextStreamID=0),r=r.toString(36),r=\"!\"+\"0000\".substr(0,3-r.length)+r;var n=E(t,e,r);return t.pendingRes[r]=function(e,t){e?n.emit(\"end\",e):t&&0!==t.length?n.emit(\"data\",t):n.emit(\"end\",null)},n},r.Handlers=M,M.prototype.handleBufferRequest=function(e,t){e?this.reqHandlers[e]=t:this.reqFallbackHandler=t},M.prototype.handleRequest=function(e,t){return this.handleBufferRequest(e,(function(e,r,n){var o=function(e){return r(JSON.stringify(e))};o.error=r.error;var s=f(e);t(s,o,n)}))},M.prototype.handleBufferNotification=function(e,t){e?this.noteHandlers[e]=t:this.noteFallbackHandler=t},M.prototype.handleNotification=function(e,t){this.handleBufferNotification(e,(function(e,r){t(f(e),r)}))},M.prototype.findRequestHandler=function(e){return this.reqHandlers[e]||this.reqFallbackHandler},M.prototype.findNotificationHandler=function(e){return this.noteHandlers[e]||this.noteFallbackHandler},void 0!==e.gotalkResponderAt){var _=e.gotalkResponderAt;_&&_.ws&&(c.defaultResponderAddress=(\"https:\"==document.location.protocol?\"wss://\":\"ws://\")+document.location.host+_.ws),delete e.gotalkResponderAt}p.prototype.open=function(e,t){if(t||\"function\"!=typeof e||(t=e,e=null),!e){if(!c.defaultResponderAddress)throw new Error(\"address not specified (responder has not announced any default address)\");e=c.defaultResponderAddress}return function(e,t,r){var n;try{(n=new WebSocket(t)).binaryType=\"arraybuffer\",n.onclose=function(e){var t=new Error(\"connection failed\");r&&r(t)},n.onopen=function(t){n.onerror=void 0,e.adoptWebSocket(n),e.handshake(),r&&r(null,e),e.emit(\"open\"),e.startReading()},n.onmessage=function(e){n._bufferedMessages||(n._bufferedMessages=[]),n._bufferedMessages.push(e.data)}}catch(t){r&&r(t),e.emit(\"close\",t)}}(this,e,t),this},c.open=function(e,t){return p(c.defaultHandlers).open(e,t)},p.prototype.openKeepAlive=function(e){var t=this;return t.keepalive&&t.keepalive.disable(),t.keepalive=d(t,e),t.keepalive.enable(),t},c.connection=function(e){return p(c.defaultHandlers).openKeepAlive(e)},c.defaultHandlers=M(),c.handleBufferRequest=function(e,t){return c.defaultHandlers.handleBufferRequest(e,t)},c.handle=function(e,t){return c.defaultHandlers.handleRequest(e,t)},c.handleBufferNotification=function(e,t){return c.defaultHandlers.handleBufferNotification(e,t)},c.handleNotification=function(e,t){return c.defaultHandlers.handleNotification(e,t)}}(t={exports:{}}),e.gotalk=t.exports}(this);"
const BrowserLibSourceMapString = "{\"version\":3,\"sources\":[\".gotalk-debug.js\"],\"names\":[\"global\",\"__mainapi\",\"__mod\",\"__api\",\"evalDepth\",\"require\",\"name\",\"m\",\"id\",\"replace\",\"i\",\"f\",\"exports\",\"module\",\"EventEmitter\",\"prototype\",\"addListener\",\"type\",\"listener\",\"TypeError\",\"this\",\"__events\",\"Object\",\"defineProperty\",\"value\",\"enumerable\",\"writable\",\"listeners\",\"undefined\",\"push\",\"on\",\"once\",\"fired\",\"trigger_event_once\",\"removeListener\",\"apply\",\"arguments\",\"p\",\"indexOf\",\"splice\",\"length\",\"removeAllListeners\",\"emit\",\"L\",\"args\",\"Array\",\"slice\",\"call\",\"console\",\"log\",\"mixin\",\"obj\",\"proto\",\"__proto__\",\"Buf\",\"Uint8Array\",\"v\",\"ArrayBuffer\",\"netAccess\",\"protocol\",\"s\",\"addr\",\"minReconnectDelay\",\"maxReconnectDelay\",\"ctx\",\"open\",\"retry\",\"openTimer\",\"opentime\",\"delay\",\"isEnabled\",\"isConnected\",\"enable\",\"enabled\",\"disable\",\"clearTimeout\",\"err\",\"Date\",\"available\",\"onLine\",\"document\",\"location\",\"hostname\",\"isGotalkProtocolError\",\"code\",\"ErrorTimeout\",\"Math\",\"min\",\"max\",\"setTimeout\",\"addEventListener\",\"create\",\"navigator\",\"ev\",\"utf8\",\"Version\",\"MsgTypeSingleReq\",\"MsgTypeStreamReq\",\"MsgTypeRetryRes\",\"MsgTypeStreamReqPart\",\"MsgTypeSingleRes\",\"MsgTypeStreamRes\",\"MsgTypeErrorRes\",\"MsgTypeNotification\",\"MsgTypeHeartbeat\",\"MsgTypeProtocolError\",\"copyBufFixnum\",\"b\",\"start\",\"n\",\"digits\",\"c\",\"y\",\"toString\",\"z\",\"isNaN\",\"charCodeAt\",\"makeBufFixnum\",\"parseIntBuf\",\"radix\",\"parseInt\",\"String\",\"fromCharCode\",\"ErrorAbnormal\",\"ErrorUnsupported\",\"ErrorInvalidMsg\",\"HeartbeatMsgMaxLoad\",\"binary\",\"makeFixnum\",\"versionBuf\",\"parseVersion\",\"parseMsg\",\"t\",\"namez\",\"wait\",\"subarray\",\"decode\",\"size\",\"makeMsg\",\"nameb\",\"encode\",\"set\",\"makeHeartbeatMsg\",\"load\",\"round\",\"getTime\",\"makeStrFixnum\",\"substr\",\"text\",\"buf\",\"substring\",\"sizeOf\",\"mask8\",\"TextDecoder\",\"decoder\",\"encoder\",\"TextEncoder\",\"lead\",\"e\",\"j\",\"txt\",\"bin\",\"keepalive\",\"gotalk\",\"decodeJSON\",\"JSON\",\"parse\",\"Sock\",\"handlers\",\"heartbeatInterval\",\"ws\",\"nextOpID\",\"nextStreamID\",\"pendingRes\",\"hasPendingRes\",\"get\",\"k\",\"pendingClose\",\"websocketCloseStatus\",\"1000\",\"1001\",\"1002\",\"1003\",\"1005\",\"1006\",\"1007\",\"1008\",\"1009\",\"adoptWebSocket\",\"readyState\",\"WebSocket\",\"OPEN\",\"Error\",\"binaryType\",\"onclose\",\"_gotalkCloseError\",\"causedByErr\",\"stopSendingHeartbeats\",\"onmessage\",\"onerror\",\"resetSock\",\"_bufferedMessages\",\"data\",\"adopt\",\"rwc\",\"handshake\",\"send\",\"end\",\"close\",\"address\",\"url\",\"ErrAbnormal\",\"ErrUnsupported\",\"ErrInvalidMsg\",\"ErrTimeout\",\"sendHeartbeat\",\"startSendingHeartbeats\",\"_sendHeartbeatsTimer\",\"startReading\",\"msg\",\"readMsg\",\"errcode\",\"readMsgPayload\",\"handleMsg\",\"closeError\",\"forEach\",\"msgHandlers\",\"handleRes\",\"payload\",\"callback\",\"msgHandler\",\"handler\",\"result\",\"findRequestHandler\",\"outbuf\",\"sendMsg\",\"error\",\"errstr\",\"message\",\"stack\",\"findNotificationHandler\",\"time\",\"payloadSize\",\"bufferRequest\",\"op\",\"bufferNotify\",\"request\",\"stringify\",\"notify\",\"Promise\",\"requestp\",\"resolve\",\"reject\",\"res\",\"bufferRequestp\",\"StreamRequest\",\"Handlers\",\"reqHandlers\",\"reqFallbackHandler\",\"noteHandlers\",\"noteFallbackHandler\",\"write\",\"ended\",\"started\",\"streamRequest\",\"req\",\"handleBufferRequest\",\"handleRequest\",\"resultWrapper\",\"handleBufferNotification\",\"handleNotification\",\"gotalkResponderAt\",\"at\",\"defaultResponderAddress\",\"host\",\"onopen\",\"openWebSocket\",\"onConnect\",\"defaultHandlers\",\"openKeepAlive\",\"connection\",\"handle\",\"__main\"],\"mappings\":\"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\"}"
