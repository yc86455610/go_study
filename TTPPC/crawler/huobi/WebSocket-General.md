
          <p>##1.引言
###1.1目的</p>
<p>本文描述了火币网实时行情开放接口API，获取火币网实时行情数据功能。可以帮助开发人员快速入门并掌握开发技能，同时也可以作为日后火币网API接口参数及参数类型的速查手册。</p>
<p>###1.2读者对象</p>
<p>预期读者为业务需求人员、开发经理、项目经理、架构设计师、开发人员、测试人员、设计评审人员、用户文档编写者。</p>
<p>###1.3接口技术介绍</p>
<p>火币网实时行情接口采用Socket.IO协议实现，兼容websocket, flashsocket, xhr-polling, jsonp-polling等链接方式。无论是普通web、webapp（html5），还是native的客户端。无论是电脑，还是移动设备，都可以轻松接入。</p>
<p>Socket.IO的版本是目前最稳定的0.9.16，如果用js语言，请用socket.io-client的0.9.16版本，其它主流语言，在github上均有Socket.IO的开源库。</p>

        