
          <h1>
<a id="user-content-错误码" class="anchor" href="#%E9%94%99%E8%AF%AF%E7%A0%81" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>错误码</h1>
<h3>
<a id="user-content-错误信息返回格式" class="anchor" href="#%E9%94%99%E8%AF%AF%E4%BF%A1%E6%81%AF%E8%BF%94%E5%9B%9E%E6%A0%BC%E5%BC%8F" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>错误信息返回格式</h3>
<pre><code>{
  &#34;id&#34;: &#34;id generate by client&#34;,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;err-code&#34;,
  &#34;err-msg&#34;: &#34;err-message&#34;,
  &#34;ts&#34;: 1487152091345
}
</code></pre>
<p>注：<code>&#34;ts&#34;</code> 为错误信息生成的时间戳，单位：毫秒</p>
<table>
<thead>
<tr>
<th>err-msg</th>
<th>描述</th>
</tr>
</thead>
<tbody>
<tr>
<td>invalid topic market.invalidsymbol.kline.1min</td>
<td>错误的 symbol</td>
</tr>
<tr>
<td>invalid topic market.btccny.kline.3min</td>
<td>错误的 topic</td>
</tr>
<tr>
<td>unsub with not subbed topic market.btccny.trade.detail</td>
<td>取消订阅一个尚未订阅的 topic</td>
</tr>
<tr>
<td>unsub with not subbed topic not-exists-topic</td>
<td>取消订阅一个不存在的 topic</td>
</tr>
<tr>
<td>invalid topic market.invalidsymbol.trade.detail</td>
<td>错误的 symbol</td>
</tr>
</tbody>
</table>

        