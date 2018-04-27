
          <h2>
<a id="user-content-交易-api-请求过程" class="anchor" href="#%E4%BA%A4%E6%98%93-api-%E8%AF%B7%E6%B1%82%E8%BF%87%E7%A8%8B" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>交易 API 请求过程</h2>
<ol>
<li>
<p>用户根据文档组织交易要素。</p>
</li>
<li>
<p>通过POST方式提交交易请求。</p>
</li>
<li>
<p>请求头信息中必须声明 Content-Type:application/x-www-form-urlencoded</p>
</li>
<li>
<p>火币网处理请求，并返回相应的JSON格式结果。</p>
</li>
<li>
<p>交易API目前只支持https请求</p>
</li>
<li>
<p>apiV3 提交地址 <a href="https://api.huobi.com/apiv3" rel="nofollow">https://api.huobi.com/apiv3</a></p>
</li>
<li>
<p>火币网API类似表单提交</p>
</li>
<li>
<p>生成sign时，参与加密的参数按照a-z排序，签名MD5必须是小写字母。</p>
</li>
<li>
<p>限制频率（每个接口，只针对交易api，行情api不限制）：withdraw_coin接口访问频率限制为1分钟10次，其余接口限制为10秒100次</p>
</li>
</ol>
<p>签名时的字符一律采用UTF-8编码。MD5散列后每个byte采用两个16进制小写字母高位在前低位在后表示。例如：</p>
<pre lang="code"><code>md5(&#34;比特币你好&#34;) = d6b6e11652b0c93c4f14cfb84c380541
md5(&#34;ABC123abc&#34;) = 85716f0702d2d464803e1366a7678d0b
</code></pre>

        