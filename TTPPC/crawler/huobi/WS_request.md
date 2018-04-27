
          <h1>
<a id="user-content-请求与订阅说明" class="anchor" href="#%E8%AF%B7%E6%B1%82%E4%B8%8E%E8%AE%A2%E9%98%85%E8%AF%B4%E6%98%8E" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>请求与订阅说明</h1>
<h3>
<a id="user-content-1-访问地址" class="anchor" href="#1-%E8%AE%BF%E9%97%AE%E5%9C%B0%E5%9D%80" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>1. 访问地址</h3>
<ul>
<li>Pro 站行情请求地址为：wss://api.huobipro.com/ws</li>
<li>HADAX 站行情请求地址为：wss://api.hadax.com/ws</li>
</ul>
<h3>
<a id="user-content-2-数据压缩" class="anchor" href="#2-%E6%95%B0%E6%8D%AE%E5%8E%8B%E7%BC%A9" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>2. 数据压缩</h3>
<p>WebSocket API 返回的所有数据都进行了 GZIP 压缩，需要 client 在收到数据之后解压，推荐使用pako。（<a href="https://github.com/nodeca/pako">【pako】</a> 是一个支持压缩和解压 GZIP 的库）</p>
<h3>
<a id="user-content-3-websocket库" class="anchor" href="#3-websocket%E5%BA%93" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>3. WebSocket库</h3>
<p><a href="https://github.com/websockets/ws">【ws】</a> 是 Node.js 下的 WebSocket 库。</p>
<h3>
<a id="user-content-4-心跳" class="anchor" href="#4-%E5%BF%83%E8%B7%B3" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>4. 心跳</h3>
<p>WebSocket API 支持双向心跳，无论是 Server 还是 Client 都可以发起 <code>ping</code> message，对方返回 <code>pong</code> message。</p>
<p>WebSocket Server 发送心跳：</p>
<pre><code>{&#34;ping&#34;: 18212558000}
</code></pre>
<p>WebSocket Client 应该返回：</p>
<pre><code> {&#34;pong&#34;: 18212558000}
</code></pre>
<p>注：返回的数据里面的 <code>&#34;pong&#34;</code> 的值为收到的 <code>&#34;ping&#34;</code> 的值
注：WebSocket Client 和 WebSocket Server 建立连接之后，WebSocket Server 每隔 <code>5s</code>（这个频率可能会变化） 会向 WebSocket Client 发起一次心跳，WebSocket Client 忽略心跳2次后，WebSocket Server 将会主动断开连接。</p>
<p>WebSocket Client 发送心跳：</p>
<pre><code>{&#34;ping&#34;: 18212553000}
</code></pre>
<p>注：发送的 message 里面，&#34;ping&#34; 的值必须为 Long 类型，否则返回错误信息：</p>
<pre><code>{
  &#34;ts&#34;: 1492420473027,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;invalid ping&#34;
}
</code></pre>
<p>WebSocket Server 会返回：</p>
<pre><code>{&#34;pong&#34;: 18212553000}
</code></pre>
<p>注：返回的数据里面的 <code>&#34;pong&#34;</code> 的值为收到的 <code>&#34;ping&#34;</code> 的值</p>
<h3>
<a id="user-content-5-topic格式" class="anchor" href="#5-topic%E6%A0%BC%E5%BC%8F" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>5. topic格式</h3>
<p>订阅数据和请求数据都要使用 <code>topic</code>，<code>topic</code> 的语法如下：</p>
<table>
<thead>
<tr>
<th>topic 类型</th>
<th>topic 语法</th>
<th>描述</th>
</tr>
</thead>
<tbody>
<tr>
<td>KLine</td>
<td>market.$symbol.kline.$period</td>
<td>$period 可选值：{ 1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year }</td>
</tr>
<tr>
<td>Market Depth</td>
<td>market.$symbol.depth.$type</td>
<td>$type 可选值：{ step0, step1, step2, step3, step4, step5 } （合并深度0-5）；step0时，不合并深度</td>
</tr>
<tr>
<td>Trade Detail</td>
<td>market.$symbol.trade.detail</td>
<td></td>
</tr>
<tr>
<td>Market Detail</td>
<td>market.$symbol.detail</td>
<td></td>
</tr>
</tbody>
</table>
<ul>
<li>
<code>$symbol</code> 为币种，可选值： { ethbtc, ltcbtc, etcbtc, bchbtc...... }</li>
<li>用户选择“合并深度”时，一定报价精度内的市场挂单将予以合并显示。合并深度仅改变显示方式，不改变实际成交价格。</li>
</ul>
<h3>
<a id="user-content-6-请求数据reqrep" class="anchor" href="#6-%E8%AF%B7%E6%B1%82%E6%95%B0%E6%8D%AEreqrep" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>6. 请求数据(req/rep)</h3>
<p>请求数据，仅返回一次数据</p>
<h4>
<a id="user-content-请求数据的格式" class="anchor" href="#%E8%AF%B7%E6%B1%82%E6%95%B0%E6%8D%AE%E7%9A%84%E6%A0%BC%E5%BC%8F" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>请求数据的格式</h4>
<pre><code>{
  &#34;req&#34;: &#34;topic to req&#34;,
  &#34;id&#34;: &#34;id generate by client&#34;
}
</code></pre>
<ul>
<li>
<code>&#34;req&#34;</code> 的值为 <strong>topic</strong> ，请参考 <code>&#34;5. topic格式&#34;</code> 的 <strong>topic 格式</strong>
</li>
</ul>
<p><strong>正确请求数据的例子</strong></p>
<pre><code>{
  &#34;req&#34;: &#34;market.btcusdt.kline.1min&#34;,
  &#34;id&#34;: &#34;id10&#34;
}
</code></pre>
<p>返回数据的例子：</p>
<pre><code>{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;rep&#34;: &#34;market.btcusdt.kline.1min&#34;,
  &#34;tick&#34;: [
    {
      &#34;amount&#34;: 1.6206,
      &#34;count&#34;:  3,
      &#34;id&#34;:     1494465840,
      &#34;open&#34;:   9887.00,
      &#34;close&#34;:  9885.00,
      &#34;low&#34;:    9885.00,
      &#34;high&#34;:   9887.00,
      &#34;vol&#34;:    16021.632026
    },
    {
      &#34;amount&#34;: 2.2124,
      &#34;count&#34;:  6,
      &#34;id&#34;:     1494465900,
      &#34;open&#34;:   9885.00,
      &#34;close&#34;:  9880.00,
      &#34;low&#34;:    9880.00,
      &#34;high&#34;:   9885.00,
      &#34;vol&#34;:    21859.023500
    }
  ]
}
</code></pre>
<p><strong>错误请求数据的例子</strong></p>
<pre><code>{
  &#34;req&#34;: &#34;market.invalidsymbo.kline.1min&#34;,
  &#34;id&#34;: &#34;id10&#34;
}
</code></pre>
<p>返回的错误信息的例子：</p>
<pre><code>{
  &#34;status&#34;: &#34;error&#34;,
  &#34;id&#34;: &#34;id10&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;invalid topic market.invalidsymbol.trade.detail&#34;,
  &#34;ts&#34;: 1494483996521
}
</code></pre>
<h3>
<a id="user-content-7-订阅数据subpub" class="anchor" href="#7-%E8%AE%A2%E9%98%85%E6%95%B0%E6%8D%AEsubpub" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>7. 订阅数据(sub/pub)</h3>
<h4>
<a id="user-content-订阅数据的格式" class="anchor" href="#%E8%AE%A2%E9%98%85%E6%95%B0%E6%8D%AE%E7%9A%84%E6%A0%BC%E5%BC%8F" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>订阅数据的格式</h4>
<p>成功建立和 WebSocket API 的连接之后，向 Server 发送如下格式的数据来订阅数据：</p>
<pre><code>{
  &#34;sub&#34;: &#34;topic to sub&#34;,
  &#34;id&#34;: &#34;id generate by client&#34;
}
</code></pre>
<p><strong>正确订阅的例子</strong></p>
<p>正确订阅：</p>
<pre><code>{
  &#34;sub&#34;: &#34;market.btcusdt.kline.1min&#34;,
  &#34;id&#34;: &#34;id1&#34;
}
</code></pre>
<ul>
<li>
<code>&#34;sub&#34;</code> 的值为 <strong>topic</strong> ，请参考 <code>&#34;5. topic格式&#34;</code> 的 <strong>topic 格式</strong>
</li>
</ul>
<p>订阅成功返回数据的例子：</p>
<pre><code>{
  &#34;id&#34;: &#34;id1&#34;,
  &#34;status&#34;: &#34;ok&#34;,
  &#34;subbed&#34;: &#34;market.btcusdt.kline.1min&#34;,
  &#34;ts&#34;: 1489474081631
}
</code></pre>
<p>之后每当 KLine 有更新时，client 会收到数据，例子：</p>
<pre><code>{
  &#34;ch&#34;: &#34;market.btcusdt.kline.1min&#34;,
  &#34;ts&#34;: 1489474082831,
  &#34;tick&#34;: {
    &#34;id&#34;: 1489464480,
    &#34;amount&#34;: 0.0,
    &#34;count&#34;: 0,
    &#34;open&#34;: 7962.62,
    &#34;close&#34;: 7962.62,
    &#34;low&#34;: 7962.62,
    &#34;high&#34;: 7962.62,
    &#34;vol&#34;: 0.0
  }
}
</code></pre>
<p>tick 说明:</p>
<pre><code>  &#34;tick&#34;: {
    &#34;id&#34;: K线id,
    &#34;amount&#34;: 成交量,
    &#34;count&#34;: 成交笔数,
    &#34;open&#34;: 开盘价,
    &#34;close&#34;: 收盘价,当K线为最晚的一根时，是最新成交价
    &#34;low&#34;: 最低价,
    &#34;high&#34;: 最高价,
    &#34;vol&#34;: 成交额, 即 sum(每一笔成交价 * 该笔的成交量)
  }

</code></pre>
<p><strong>错误订阅的例子</strong></p>
<p>错误订阅（错误的 symbol）：</p>
<pre><code>{
  &#34;sub&#34;: &#34;market.invalidsymbol.kline.1min&#34;,
  &#34;id&#34;: &#34;id2&#34;
}
</code></pre>
<p>订阅失败返回数据的例子：</p>
<pre><code>{
  &#34;id&#34;: &#34;id2&#34;,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;invalid topic market.invalidsymbol.kline.1min&#34;,
  &#34;ts&#34;: 1494301904959
}
</code></pre>
<p>错误订阅（错误的 topic）：</p>
<pre><code>{
  &#34;sub&#34;: &#34;market.btcusdt.kline.3min&#34;,
  &#34;id&#34;: &#34;id3&#34;
}
</code></pre>
<p>订阅失败返回数据的例子：</p>
<pre><code>{
  &#34;id&#34;: &#34;id3&#34;,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;invalid topic market.btcusdt.kline.3min&#34;,
  &#34;ts&#34;: 1494310283622
}
</code></pre>
<h3>
<a id="user-content-8-取消订阅unsub" class="anchor" href="#8-%E5%8F%96%E6%B6%88%E8%AE%A2%E9%98%85unsub" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>8. 取消订阅(unsub)</h3>
<h4>
<a id="user-content-取消订阅的格式" class="anchor" href="#%E5%8F%96%E6%B6%88%E8%AE%A2%E9%98%85%E7%9A%84%E6%A0%BC%E5%BC%8F" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>取消订阅的格式</h4>
<p>WebSocket Client 订阅数据之后，可以取消订阅，取消订阅之后 WebSocket Server 将不会再发送该 topic 的数据，取消订阅的格式如下：</p>
<pre><code>{
  &#34;unsub&#34;: &#34;topic to unsub&#34;,
  &#34;id&#34;: &#34;id generate by client&#34;
}
</code></pre>
<p><strong>正确取消订阅的例子</strong></p>
<p>正确取消订阅的例子：</p>
<pre><code>{
  &#34;unsub&#34;: &#34;market.btcusdt.trade.detail&#34;,
  &#34;id&#34;: &#34;id4&#34;
}
</code></pre>
<p>取消订阅成功返回信息的例子：</p>
<pre><code>{
  &#34;id&#34;: &#34;id4&#34;,
  &#34;status&#34;: &#34;ok&#34;,
  &#34;unsubbed&#34;: &#34;market.btcusdt.trade.detail&#34;,
  &#34;ts&#34; 1494326028889
}
</code></pre>
<p><strong>错误取消订阅的例子</strong></p>
<p>错误取消订阅的例子（取消订阅一个尚未订阅的 topic）：</p>
<pre><code>{
  &#34;unsub&#34;: &#34;market.btcusdt.trade.detail&#34;,
  &#34;id&#34;: &#34;id5&#34;
}
</code></pre>
<p>返回的错误信息的例子</p>
<pre><code>{
  &#34;id&#34;: &#34;id5&#34;,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;unsub with not subbed topic market.btcusdt.trade.detail&#34;,
  &#34;ts&#34;: 1494326217428
}
</code></pre>
<p>错误取消订阅的例子（取消订阅一个不存在的 topic）：</p>
<pre><code>{
  &#34;unsub&#34;: &#34;not-exists-topic&#34;,
  &#34;id&#34;: &#34;id5&#34;
}
</code></pre>
<p>返回的错误信息的例子：</p>
<pre><code>{
  &#34;id&#34;: &#34;id5&#34;,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;unsub with not subbed topic not-exists-topic&#34;,
  &#34;ts&#34;: 1494326318809
}
</code></pre>

        