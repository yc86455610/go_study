
          <h2>
<a id="user-content-安全认证" class="anchor" href="#%E5%AE%89%E5%85%A8%E8%AE%A4%E8%AF%81" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>安全认证</h2>
<p>目前关于apikey申请和修改，请在“账户 - API管理”页面进行相关操作。其中AccessKey为API 访问密钥，SecretKey为用户对请求进行签名的密钥（仅申请时可见）。Pro站和HADAX站apikey通用。</p>
<ul>
<li><strong>重要提示：这两个密钥与账号安全紧密相关，无论何时都请勿向其它人透露。</strong></li>
</ul>
<h2>
<a id="user-content-合法请求结构" class="anchor" href="#%E5%90%88%E6%B3%95%E8%AF%B7%E6%B1%82%E7%BB%93%E6%9E%84" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>合法请求结构</h2>
<p>基于安全考虑，除行情API 外的 API 请求都必须进行签名运算。一个合法的请求由以下几部分组成：</p>
<ul>
<li>
<p>方法请求地址
即访问服务器地址：api.huobi.pro或者api.hadax.com后面跟上方法名，比如api.huobi.pro/v1/order/orders。</p>
</li>
<li>
<p>API 访问密钥（AccessKeyId）
您申请的 APIKEY 中的AccessKey。</p>
</li>
<li>
<p>签名方法（SignatureMethod）
用户计算签名的基于哈希的协议，此处使用 HmacSHA256。</p>
</li>
<li>
<p>签名版本（SignatureVersion）
签名协议的版本，此处使用2。</p>
</li>
<li>
<p>时间戳（Timestamp）
您发出请求的时间 <strong>(UTC 时区)</strong>  <strong>(UTC 时区)</strong>  <strong>(UTC 时区)</strong> 。在查询请求中包含此值有助于防止第三方截取您的请求。如：2017-05-11T16:22:06。再次强调是 <strong>(UTC 时区)</strong> 。</p>
</li>
<li>
<p>必选和可选参数
每个方法都有一组用于定义 API 调用的必需参数和可选参数。可以在每个方法的说明中查看这些参数及其含义。
请一定注意：对于<strong>GET</strong>请求，每个方法自带的参数都需要进行签名运算；
对于<strong>POST</strong>请求，每个方法自带的参数不进行签名认证，即<strong>POST</strong>请求中需要进行签名运算的只有AccessKeyId、SignatureMethod、SignatureVersion、Timestamp四个参数，其它参数放在body中。</p>
</li>
<li>
<p>签名
签名计算得出的值，用于确保签名有效和未被篡改。</p>
</li>
</ul>
<p>例：</p>
<pre><code>https://api.huobi.pro/v1/order/orders?
AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx
&amp;SignatureMethod=HmacSHA256
&amp;SignatureVersion=2
&amp;Timestamp=2017-05-11T15%3A19%3A30
&amp;order-id=1234567890
&amp;Signature=calculated value
</code></pre>
<h2>
<a id="user-content-签名运算" class="anchor" href="#%E7%AD%BE%E5%90%8D%E8%BF%90%E7%AE%97" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>签名运算</h2>
<p>API 请求在通过 Internet 发送的过程中极有可能被篡改。为了确保请求未被更改，我们会要求用户在每个请求中带上签名（行情 API 除外），来校验参数或参数值在传输途中是否发生了更改。</p>
<p><strong>计算签名所需的步骤</strong>：</p>
<ol>
<li>规范要计算签名的请求</li>
</ol>
<p>因为使用 HMAC 进行签名计算时，使用不同内容计算得到的结果会完全不同。所以在进行签名计算前，请先对请求进行规范化处理。下面以查询某订单详情请求为例进行说明</p>
<pre><code>https://api.huobi.pro/v1/order/orders?
AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx
&amp;SignatureMethod=HmacSHA256
&amp;SignatureVersion=2
&amp;Timestamp=2017-05-11T15:19:30
&amp;order-id=1234567890
</code></pre>
<ol start="2">
<li>请求方法（GET 或 POST），后面添加换行符\n。</li>
</ol>
<pre><code>GET\n
</code></pre>
<ol start="3">
<li>添加小写的访问地址，后面添加换行符\n。</li>
</ol>
<pre><code>api.huobi.pro\n
</code></pre>
<ol start="4">
<li>访问方法的路径，后面添加换行符\n。</li>
</ol>
<pre><code>/v1/order/orders\n
</code></pre>
<ol start="5">
<li>按照<strong>ASCII码的顺序对参数名进行排序</strong>(<strong>使用 UTF-8 编码</strong>，<strong>且进行了 URI 编码</strong>，<strong>十六进制字符必须大写</strong>，如‘:’会被编码为&#39;%3A&#39;，空格被编码为&#39;%20&#39;)。</li>
</ol>
<p>例如，下面是请求参数的原始顺序，进行过编码后。</p>
<pre><code>AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx
order-id=1234567890
SignatureMethod=HmacSHA256
SignatureVersion=2
Timestamp=2017-05-11T15%3A19%3A30
</code></pre>
<p>这些参数会被排序为：</p>
<pre><code>AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx
SignatureMethod=HmacSHA256
SignatureVersion=2
Timestamp=2017-05-11T15%3A19%3A30
order-id=1234567890
</code></pre>
<p>按照以上顺序，将各参数使用字符’&amp;’连接。</p>
<pre><code>AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx&amp;SignatureMethod=HmacSHA256&amp;SignatureVersion=2&amp;Timestamp=2017-05-11T15%3A19%3A30&amp;order-id=1234567890
</code></pre>
<p>组成最终的要进行签名计算的字符串如下：</p>
<pre><code>GET\n
api.huobi.pro\n
/v1/order/orders\n
AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx&amp;SignatureMethod=HmacSHA256&amp;SignatureVersion=2&amp;Timestamp=2017-05-11T15%3A19%3A30&amp;order-id=1234567890
</code></pre>
<ol start="6">
<li>计算签名，将以下两个参数传入加密哈希函数：</li>
</ol>
<ul>
<li>要进行签名计算的字符串</li>
</ul>
<pre><code>GET\n
api.huobi.pro\n
/v1/order/orders\n
AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx&amp;SignatureMethod=HmacSHA256&amp;SignatureVersion=2&amp;Timestamp=2017-05-11T15%3A19%3A30&amp;order-id=1234567890
</code></pre>
<ul>
<li>进行签名的密钥（SecretKey）</li>
</ul>
<pre><code>b0xxxxxx-c6xxxxxx-94xxxxxx-dxxxx
</code></pre>
<p>得到签名计算结果并进行 Base64编码</p>
<pre><code>4F65x5A2bLyMWVQj3Aqp+B4w+ivaA7n5Oi2SuYtCJ9o=
</code></pre>
<ol start="7">
<li>
<p>将上述值作为参数Signature的取值添加到 API 请求中。
将此参数添加到请求时，<strong>必须将该值进行 URI 编码</strong>。</p>
</li>
<li>
<p>最终，发送到服务器的 API 请求应该为：</p>
</li>
</ol>
<pre><code>https://api.huobi.pro/v1/order/orders?AccessKeyId=e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx&amp;order-id=1234567890&amp;SignatureMethod=HmacSHA256&amp;SignatureVersion=2&amp;Timestamp=2017-05-11T15%3A19%3A30&amp;Signature=4F65x5A2bLyMWVQj3Aqp%2BB4w%2BivaA7n5Oi2SuYtCJ9o%3D
</code></pre>

        