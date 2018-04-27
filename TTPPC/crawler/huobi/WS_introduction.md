
          <h1>
<a id="user-content-websocket-api简介" class="anchor" href="#websocket-api%E7%AE%80%E4%BB%8B" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>WebSocket API简介</h1>
<p>WebSocket协议是基于TCP的一种新的网络协议。它实现了客户端与服务器之间在单个 tcp 连接上的全双工通信，由服务器主动发送信息给客户端，减少了频繁的身份验证等不必要的开销。其最大优点有两个：</p>
<ul>
<li>
<p>两方请求的 header 数据很小，大概只有2 Bytes。</p>
</li>
<li>
<p>服务器不再是被动的接到客户端的请求后才返回数据，而是有了新数据后主动推送给客户端。</p>
</li>
</ul>
<p>以上 WebSocket 协议带来的优点使得其十分适用于数字货币行情和交易这种实时性强的接口。</p>
<p>WebSocket 只支持行情查询，交易接口将在后续提供。在使用中如果遇到问题，请加技术讨论 QQ 群: 火币网API交流群(3) 597821383，我们将尽力帮您答疑解惑（加群时请注明uid（账户管理中）和编程语言）。</p>

        