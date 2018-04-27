
          <h2>
<a id="user-content-请求说明" class="anchor" href="#%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>请求说明</h2>
<ol>
<li>访问地址</li>
</ol>
<p>Pro 站：
行情： <a href="https://api.huobipro.com/market" rel="nofollow">https://api.huobipro.com/market</a>
交易： <a href="https://api.huobipro.com/v1" rel="nofollow">https://api.huobipro.com/v1</a></p>
<p>HADAX 站：
行情： <a href="https://api.hadax.com/market" rel="nofollow">https://api.hadax.com/market</a>
交易： <a href="https://api.hadax.com/v1" rel="nofollow">https://api.hadax.com/v1</a></p>
<ol start="2">
<li>POST请求头信息中必须声明 Content-Type:application/json;GET请求头信息中必须声明 Content-Type:application/x-www-form-urlencoded。(汉语用户建议设置 Accept-Language:zh-cn)</li>
<li>所有请求参数请按照 API 说明进行参数封装。</li>
<li>将封装好参数的 API 请求通过 POST 或 GET 的方式提交到服务器。</li>
<li>火币网处理请求，并返回相应的 JSON 格式结果。</li>
<li>请使用 https 请求。</li>
<li>限制频率（每个接口，只针对交易api，行情api不限制）为10秒100次。</li>
<li>查询资产详情方法调用顺序：查询当前用户的所有账户-&gt;查询指定账户的余额</li>
<li>支持所有Pro站上交易中的交易对，上新币保持与网站同步。</li>
</ol>

        