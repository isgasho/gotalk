!function(e){"use strict";var t,r={},n={},o=0,s=function(e){var t,s=e.replace(/^.\//,"");if(!(t=n[s])){for(var a=0;a<o;a++)". ";var i=r[s];i&&o<100&&(r[s]=null,n[s]={exports:{}},++o,i(n[s]),--o,n[s]=n[s].exports),t=n[s]}return t};r.EventEmitter=function(e){e.exports;function t(){}e.exports=t,t.prototype.addListener=function(e,t){if("function"!=typeof t)throw TypeError("listener must be a function");if(!this.__events)return Object.defineProperty(this,"__events",{value:{},enumerable:!1,writable:!0}),this.__events[e]=[t],this;var r=this.__events[e];return void 0===r?(this.__events[e]=[t],this):(r.push(t),this)},t.prototype.on=t.prototype.addListener,t.prototype.once=function(e,t){var r=!1,n=function(){this.removeListener(e,n),r||(r=!0,t.apply(this,arguments))};return this.on(e,n)},t.prototype.removeListener=function(e,t){var r,n=this.__events?this.__events[e]:void 0;if(void 0!==n){for(;-1!==(r=n.indexOf(t));)n.splice(r,1);return 0===n.length&&delete this.__events[e],n.length}return this},t.prototype.removeAllListeners=function(e){this.__events&&(e?delete this.__events[e]:delete this.__events)},t.prototype.listeners=function(e){return e?this.__events?this.__events[e]:void 0:this.__events},t.prototype.emit=function(e){var t=this.__events?this.__events[e]:void 0;if(void 0===t)return!1;for(var r=0,n=t.length,o=Array.prototype.slice.call(arguments,1);r!==n;++r)t[r]||console.log("e",e,r,o),t[r].apply(this,o);return!0},t.mixin=function(e){for(var r=e;r;){if(r.__proto__===Object.prototype)return r.__proto__=t.prototype,e;r=r.__proto__}return e}},r.buf=function(e){var t;e.exports;if("undefined"!=typeof Uint8Array){s("./utf8");t=function(e){return e instanceof Uint8Array?e:new Uint8Array(e instanceof ArrayBuffer?e:new ArrayBuffer(e))}}e.exports=t},r.keepalive=function(e){e.exports;var t=s("./netaccess"),r=s("./protocol");e.exports=function(e,n,o,s){o?o<100&&(o=100):o=500,(!s||s<o)&&(s=5e3);var a,i,u,l,d,c=0;return a={isEnabled:!1,isConnected:!1,enable:function(){a.enabled||(a.enabled=!0,c=0,a.isConnected||i())},disable:function(){a.enabled&&(clearTimeout(l),a.enabled=!1,c=0)}},i=function(){clearTimeout(l),e.open(n,(function(t){d=new Date,t?u(t):(c=0,a.isConnected=!0,e.once("close",u))}))},u=function(e){if(clearTimeout(l),a.isConnected=!1,a.enabled){if(t.available&&!t.onLine&&("undefined"==typeof document||!document.location||"localhost"===document.location.hostname||"127.0.0.1"===document.location.hostname||"[::1]"===document.location.hostname))return t.once("online",u),void(c=0);c=e?e.isGotalkProtocolError?e.code===r.ErrorTimeout?0:s:c?Math.min(s,2*c):o:Math.max(0,o-(new Date-d)),l=setTimeout(i,c)}},a}},r.netaccess=function(t){t.exports;var r,n=s("./EventEmitter");void 0!==e&&e.addEventListener?(r=Object.create(n.prototype,{available:{value:!0,enumerable:!0},onLine:{value:!0,enumerable:!0,writable:!0}}),"undefined"!=typeof navigator&&(r.onLine=navigator.onLine),e.addEventListener("offline",(function(e){r.onLine=!1,r.emit("offline")})),e.addEventListener("online",(function(e){r.onLine=!0,r.emit("online")}))):r={available:!1,onLine:!0},t.exports=r},r.protocol=function(e){var t=e.exports,r=s("./buf"),n=s("./utf8");t.Version=1;var o=t.MsgTypeSingleReq=114,a=t.MsgTypeStreamReq=115,i=(t.MsgTypeStreamReqPart=112,t.MsgTypeSingleRes=82,t.MsgTypeStreamRes=83,t.MsgTypeErrorRes=69,t.MsgTypeRetryRes=101),u=t.MsgTypeNotification=110,l=t.MsgTypeHeartbeat=104,d=t.MsgTypeProtocolError=102;function c(e,t,r,n){for(var o,s=t||0,a=0,i=r.toString(16),u=n-i.length;u--;)e[s++]=48;for(;!isNaN(o=i.charCodeAt(a++));)e[s++]=o}function f(e,t){var n=r(t);return c(n,0,e,t),n}function p(e,t){return parseInt(String.fromCharCode(...e),t)}t.ErrorAbnormal=0,t.ErrorUnsupported=1,t.ErrorInvalidMsg=2,t.ErrorTimeout=3,t.HeartbeatMsgMaxLoad=65535,t.binary={makeFixnum:f,versionBuf:f(t.Version,2),parseVersion:function(e){return p(e,16)},parseMsg:function(e){var t,r,s,c,f,h=0;return f=1,(t=e[0])===l?(h=p(e.subarray(f,f+4),16),f+=4):t!==u&&t!==d&&(r=e.subarray(f,f+4),f+=4),t==o||t==a||t==u?(c=p(e.subarray(f,f+3),16),f+=3,s=n.decode(e.subarray(f,f+c)),f+=c):t===i&&(h=p(e.subarray(f,f+8),16),f+=8),{t,id:r,name:s,wait:h,size:p(e.subarray(f,f+8),16)}},makeMsg:function(e,t,o,s,a){var u,l,d=t?13:9;return o&&0!==o.length&&(d+=3+(l=n.encode(o)).length),(u=r(d))[0]=e,d=1,t&&4===t.length&&("string"==typeof t?(u[1]=t.charCodeAt(0),u[2]=t.charCodeAt(1),u[3]=t.charCodeAt(2),u[4]=t.charCodeAt(3)):(u[1]=t[0],u[2]=t[1],u[3]=t[2],u[4]=t[3]),d+=4),l&&(c(u,d,l.length,3),d+=3,u.set(l,d),d+=l.length),e===i&&(c(u,d,s,8),d+=8),c(u,d,a,8),u},makeHeartbeatMsg:function(e){var t=r(13),n=1;return t[0]=l,c(t,n,e,4),c(t,n+=4,Math.round((new Date).getTime()/1e3),8),n+=8,t}};function h(e,t){var r=e.toString(16);return"00000000".substr(0,t-r.length)+r}t.text={makeFixnum:h,versionBuf:h(t.Version,2),parseVersion:function(e){return parseInt(e.substr(0,2),16)},parseMsg:function(e){var t,r,n,s,c=0;return s=1,(t=e.charCodeAt(0))===l?(c=parseInt(e.substr(s,4),16),s+=4):t!==u&&t!==d&&(r=e.substr(s,4),s+=4),t==o||t==a||t==u?n=e.substring(s+3,e.length-8):t==i&&(c=parseInt(e.substr(s,8),16),s+=8),{t,id:r,name:n,wait:c,size:parseInt(e.substr(e.length-8),16)}},makeMsg:function(e,t,r,o,s){var a=String.fromCharCode(e);return t&&4===t.length&&(a+=t),r&&0!==r.length&&(a+=h(n.sizeOf(r),3),a+=r),e===i&&(a+=h(o,8)),a+=h(s,8)},makeHeartbeatMsg:function(e){var t=String.fromCharCode(l);return t+=h(e,4),t+=h(Math.round((new Date).getTime()/1e3),8)}}},r.utf8=function(e){var t=e.exports;function r(e){for(var t,r=0,n=0;t=e.charCodeAt(n++);r+=t>>11?3:t>>7?2:1);return r}function n(e){return 255&e}if(t.sizeOf=r,"undefined"!=typeof TextDecoder){var o=new TextDecoder("utf8"),s=new TextEncoder("utf8");t.decode=function(e){return o.decode(e)},t.encode=function(e){return s.encode(e)}}else t.decode=function(e){var t,r,o=0,s=e.length-o,a="";for(o=0;o<s;)(r=n(t=e[o++]))<128||(r>>5==6?t=(t<<6&2047)+(63&e[o++]):r>>4==14?(t=(t<<12&65535)+(n(e[o++])<<6&4095),t+=63&e[o++]):r>>3==30&&(t=(t<<18&2097151)+(n(e[o++])<<12&262143),t+=n(e[o++])<<6&4095,t+=63&e[o++])),a+=String.fromCharCode(t);return a},t.encode=function(e){for(var t,n=0,o=e.length,s=0,a=new Uint8Array(r(e));n!==o;)(t=e.charCodeAt(n++))<128?a[s++]=t:t<2048?(a[s++]=t>>6|192,a[s++]=63&t|128):t<65536?(a[s++]=t>>12|224,a[s++]=t>>6&63|128,a[s++]=63&t|128):(a[s++]=t>>18|240,a[s++]=t>>12&63|128,a[s++]=t>>6&63|128,a[s++]=63&t|128);return a}},function(t){var r=t.exports,n=s("./protocol"),o=n.text,a=n.binary,i=s("./buf"),u=s("./utf8"),l=s("./EventEmitter"),d=s("./keepalive"),c=r;function f(e){try{return JSON.parse(e)}catch(e){}}function p(e){return Object.create(p.prototype,{handlers:{value:e,enumerable:!0},protocol:{value:i?n.binary:n.text,enumerable:!0,writable:!0},heartbeatInterval:{value:2e4,enumerable:!0,writable:!0},ws:{value:null,writable:!0},keepalive:{value:null,writable:!0},nextOpID:{value:0,writable:!0},nextStreamID:{value:0,writable:!0},pendingRes:{value:{},writable:!0},hasPendingRes:{get:function(){for(var e in this.pendingRes)return!0}},pendingClose:{value:!1,writable:!0}})}c.protocol=n,c.Buf=i,p.prototype=l.mixin(p.prototype),r.Sock=p;var h={1e3:"normal",1001:"going away",1002:"protocol error",1003:"unsupported",1005:"no status",1006:"abnormal",1007:"inconsistent",1008:"invalid message",1009:"too large"};p.prototype.adoptWebSocket=function(e){var t=this;if(e.readyState!==WebSocket.OPEN)throw new Error("web socket readyState != OPEN");e.binaryType="arraybuffer",t.ws=e,e.onclose=function(r){var n=e._gotalkCloseError;n||1e3===r.code||(n=new Error("websocket closed: "+(h[r.code]||"#"+r.code))),function(e,t){if(e.pendingClose=!1,e.stopSendingHeartbeats(),e.ws&&(e.ws.onmessage=null,e.ws.onerror=null,e.ws.onclose=null,e.ws=null),e.nextOpID=0,e.hasPendingRes){var r=t||new Error("connection closed");for(var n in e.pendingRes)e.pendingRes[n](r);e.pendingRes={}}}(t,n),t.emit("close",n)},e.onmessage=function(t){e._bufferedMessages||(e._bufferedMessages=[]),e._bufferedMessages.push(t.data)}},p.prototype.adopt=function(e){if(adopt instanceof WebSocket)return this.adoptWebSocket(e);throw new Error("unsupported transport")},p.prototype.handshake=function(){this.ws.send(this.protocol.versionBuf)},p.prototype.end=function(){var e=this;e.keepalive&&(e.keepalive.disable(),e.keepalive=null),!e.pendingClose&&e.hasPendingRes?e.pendingClose=!0:e.ws&&e.ws.close(1e3)},p.prototype.address=function(){return this.ws?this.ws.url:null};var g=r.ErrAbnormal=Error("unsupported protocol");g.isGotalkProtocolError=!0,g.code=n.ErrorAbnormal;var v=r.ErrUnsupported=Error("unsupported protocol");v.isGotalkProtocolError=!0,v.code=n.ErrorUnsupported;var y=r.ErrInvalidMsg=Error("invalid protocol message");y.isGotalkProtocolError=!0,y.code=n.ErrorInvalidMsg;var b=r.ErrTimeout=Error("timeout");b.isGotalkProtocolError=!0,b.code=n.ErrorTimeout,p.prototype.sendHeartbeat=function(e){var t=this.protocol.makeHeartbeatMsg(Math.round(e*n.HeartbeatMsgMaxLoad));try{this.ws.send(t)}catch(e){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(e=new Error("socket is closed")),e}},p.prototype.startSendingHeartbeats=function(){var e=this;if(e.heartbeatInterval<10)throw new Error("Sock.heartbeatInterval is too low");clearTimeout(e._sendHeartbeatsTimer);var t=function(){clearTimeout(e._sendHeartbeatsTimer),e.sendHeartbeat(0),e._sendHeartbeatsTimer=setTimeout(t,e.heartbeatInterval)};e._sendHeartbeatsTimer=setTimeout(t,1)},p.prototype.stopSendingHeartbeats=function(){clearTimeout(this._sendHeartbeatsTimer)},p.prototype.startReading=function(){var e,t=this,r=t.ws;function s(s){if((e="string"==typeof s.data?o.parseMsg(s.data):a.parseMsg(i(s.data))).t===n.MsgTypeProtocolError){var l=e.size;l===n.ErrorAbnormal?r._gotalkCloseError=g:l===n.ErrorUnsupported?r._gotalkCloseError=v:l===n.ErrorTimeout?r._gotalkCloseError=b:r._gotalkCloseError=y,r.close(4e3+l)}else 0!==e.size&&e.t!==n.MsgTypeHeartbeat?r.onmessage=u:(t.handleMsg(e),e=null)}function u(n){var o=n.data;r.onmessage=s,t.handleMsg(e,"string"==typeof o?o:i(o)),e=null}r.onmessage=function(e){("string"==typeof e.data?o.parseVersion(e.data):a.parseVersion(i(e.data)))!==n.Version?(r._gotalkCloseError=v,t.closeError(n.ErrorUnsupported)):(r.onmessage=s,t.heartbeatInterval>0&&t.startSendingHeartbeats())},r._bufferedMessages&&(r._bufferedMessages.forEach((function(e){r.onmessage({data:e})})),r._bufferedMessages=null)};var m={};function w(e,t){var r=e.id;"string"!=typeof r&&(r=String.fromCharCode(...r));var o=this,s=o.pendingRes[r];e.t===n.MsgTypeStreamRes&&t&&0!==(t.length||t.size)||(delete o.pendingRes[r],o.pendingClose&&!o.hasPendingRes&&o.end()),"function"==typeof s&&(e.t===n.MsgTypeErrorRes?("string"!=typeof t&&(t=u.decode(t)),s(new Error(t),null)):s(null,t))}p.prototype.handleMsg=function(e,t){var r=this,o=m[e.t];o?o.call(r,e,t):(r.ws&&(r.ws._gotalkCloseError=y),r.closeError(n.ErrorInvalidMsg))},m[n.MsgTypeSingleReq]=function(e,t){var r,o,s=this;if(r=s.handlers.findRequestHandler(e.name),(o=function(t){s.sendMsg(n.MsgTypeSingleRes,e.id,null,0,t)}).error=function(t){var r=t.message||String(t);s.sendMsg(n.MsgTypeErrorRes,e.id,null,0,r)},"function"!=typeof r)o.error('no such operation "'+e.name+'"');else try{r(t,o,e.name)}catch(e){"undefined"!=typeof console&&console.error(e.stack||e),o.error("internal error")}},m[n.MsgTypeSingleRes]=w,m[n.MsgTypeStreamRes]=w,m[n.MsgTypeErrorRes]=w,m[n.MsgTypeNotification]=function(e,t){var r=this.handlers.findNotificationHandler(e.name);r&&r(t,e.name)},m[n.MsgTypeHeartbeat]=function(e){this.emit("heartbeat",{time:new Date(1e3*e.size),load:e.wait})},p.prototype.sendMsg=function(e,t,r,o,s){var a=s&&"string"==typeof s&&this.protocol===n.binary?u.sizeOf(s):s?s.length||s.size:0,i=this,l=i.protocol.makeMsg(e,t,r,o,a);try{i.ws.send(l),0!==a&&i.ws.send(s)}catch(e){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(e=new Error("socket is closed")),e}},p.prototype.closeError=function(e){var t=this;if(t.ws){try{t.ws.send(t.protocol.makeMsg(n.MsgTypeProtocolError,null,null,0,e))}catch(e){}t.ws.close(4e3+e)}};p.prototype.bufferRequest=function(e,t,r){var o=this,s=o.nextOpID++;1679616===o.nextOpID&&(o.nextOpID=0),s=s.toString(36),s="0000".substr(0,4-s.length)+s,o.pendingRes[s]=r;try{o.sendMsg(n.MsgTypeSingleReq,s,e,0,t)}catch(e){delete o.pendingRes[s],r(e)}},p.prototype.bufferNotify=function(e,t){this.sendMsg(n.MsgTypeNotification,null,e,0,t)},p.prototype.request=function(e,t,r){var n;return r?n=JSON.stringify(t):r=t,this.bufferRequest(e,n,(function(e,t){var n=f(t);return r(e,n)}))},p.prototype.notify=function(e,t){var r=JSON.stringify(t);return this.bufferNotify(e,r)},"undefined"!=typeof Promise&&(p.prototype.requestp=function(e,t){return new Promise((r,n)=>{this.request(e,t,(e,t)=>e?n(e):r(t))})},p.prototype.bufferRequestp=function(e,t){return new Promise((r,n)=>{this.bufferRequest(e,t,(e,t)=>e?n(e):r(t))})});var E=function(e,t,r){return Object.create(E.prototype,{s:{value:e},op:{value:t,enumerable:!0},id:{value:r,enumerable:!0}})};function M(){return Object.create(M.prototype,{reqHandlers:{value:{}},reqFallbackHandler:{value:null,writable:!0},noteHandlers:{value:{}},noteFallbackHandler:{value:null,writable:!0}})}if(l.mixin(E.prototype),E.prototype.write=function(e){this.ended||(this.started?this.s.sendMsg(n.MsgTypeStreamReqPart,this.id,null,0,e):(this.started=!0,this.s.sendMsg(n.MsgTypeStreamReq,this.id,this.op,0,e)),e&&0!==e.length&&0!==e.size||(this.ended=!0))},E.prototype.end=function(){this.write(null)},p.prototype.streamRequest=function(e){var t=this,r=t.nextStreamID++;46656===t.nextStreamID&&(t.nextStreamID=0),r=r.toString(36),r="!"+"0000".substr(0,3-r.length)+r;var n=E(t,e,r);return t.pendingRes[r]=function(e,t){e?n.emit("end",e):t&&0!==t.length?n.emit("data",t):n.emit("end",null)},n},r.Handlers=M,M.prototype.handleBufferRequest=function(e,t){e?this.reqHandlers[e]=t:this.reqFallbackHandler=t},M.prototype.handleRequest=function(e,t){return this.handleBufferRequest(e,(function(e,r,n){var o=function(e){return r(JSON.stringify(e))};o.error=r.error;var s=f(e);t(s,o,n)}))},M.prototype.handleBufferNotification=function(e,t){e?this.noteHandlers[e]=t:this.noteFallbackHandler=t},M.prototype.handleNotification=function(e,t){this.handleBufferNotification(e,(function(e,r){t(f(e),r)}))},M.prototype.findRequestHandler=function(e){return this.reqHandlers[e]||this.reqFallbackHandler},M.prototype.findNotificationHandler=function(e){return this.noteHandlers[e]||this.noteFallbackHandler},void 0!==e.gotalkResponderAt){var _=e.gotalkResponderAt;_&&_.ws&&(c.defaultResponderAddress=("https:"==document.location.protocol?"wss://":"ws://")+document.location.host+_.ws),delete e.gotalkResponderAt}p.prototype.open=function(e,t){if(t||"function"!=typeof e||(t=e,e=null),!e){if(!c.defaultResponderAddress)throw new Error("address not specified (responder has not announced any default address)");e=c.defaultResponderAddress}return function(e,t,r){var n;try{(n=new WebSocket(t)).binaryType="arraybuffer",n.onclose=function(e){var t=new Error("connection failed");r&&r(t)},n.onopen=function(t){n.onerror=void 0,e.adoptWebSocket(n),e.handshake(),r&&r(null,e),e.emit("open"),e.startReading()},n.onmessage=function(e){n._bufferedMessages||(n._bufferedMessages=[]),n._bufferedMessages.push(e.data)}}catch(t){r&&r(t),e.emit("close",t)}}(this,e,t),this},c.open=function(e,t){return p(c.defaultHandlers).open(e,t)},p.prototype.openKeepAlive=function(e){var t=this;return t.keepalive&&t.keepalive.disable(),t.keepalive=d(t,e),t.keepalive.enable(),t},c.connection=function(e){return p(c.defaultHandlers).openKeepAlive(e)},c.defaultHandlers=M(),c.handleBufferRequest=function(e,t){return c.defaultHandlers.handleBufferRequest(e,t)},c.handle=function(e,t){return c.defaultHandlers.handleRequest(e,t)},c.handleBufferNotification=function(e,t){return c.defaultHandlers.handleBufferNotification(e,t)},c.handleNotification=function(e,t){return c.defaultHandlers.handleNotification(e,t)}}(t={exports:{}}),e.gotalk=t.exports}(this);