
          <h2>
<a id="user-content-代码示例" class="anchor" href="#%E4%BB%A3%E7%A0%81%E7%A4%BA%E4%BE%8B" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>代码示例</h2>
<p>火币目前提供了JAVA、PHP、Python版本的代码示例，下载地址：<a href="https://github.com/huobiapi?tab=repositories">https://github.com/huobiapi?tab=repositories</a><br/>
其他语言会根据需要相继支持。在使用过程中有任何问题请联系我们的API技术讨论QQ群，我们将在第一时间帮您解决技术问题。</p>
<p>以下是PHP代码示例</p>
<div class="highlight highlight-text-html-php"><pre><span class="pl-pse">&lt;?php</span><span class="pl-s1"></span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-c"><span class="pl-c">//</span> 密匙对</span></span>
<span class="pl-s1"><span class="pl-c1">define</span>(<span class="pl-s"><span class="pl-pds">&#39;</span>ACCESS_KEY<span class="pl-pds">&#39;</span></span>, <span class="pl-s"><span class="pl-pds">&#39;</span>xxxxxxxx-xxxxxxxx-xxxxxxxx-xxxxx<span class="pl-pds">&#39;</span></span>);</span>
<span class="pl-s1"><span class="pl-c1">define</span>(<span class="pl-s"><span class="pl-pds">&#39;</span>SECRET_KEY<span class="pl-pds">&#39;</span></span>, <span class="pl-s"><span class="pl-pds">&#39;</span>xxxxxxxx-xxxxxxxx-xxxxxxxx-xxxxx<span class="pl-pds">&#39;</span></span>);</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-c"><span class="pl-c">//</span> 使用 apiv3</span></span>
<span class="pl-s1"><span class="pl-c1">define</span>(<span class="pl-s"><span class="pl-pds">&#39;</span>API_URL<span class="pl-pds">&#39;</span></span>, <span class="pl-s"><span class="pl-pds">&#39;</span>https://api.huobi.com/apiv3<span class="pl-pds">&#39;</span></span>);</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-k">function</span> <span class="pl-en">httpRequest</span>(<span class="pl-smi">$pUrl</span>, <span class="pl-smi">$pData</span>){</span>
<span class="pl-s1">	<span class="pl-smi">$tCh</span> <span class="pl-k">=</span> <span class="pl-c1">curl_init</span>();</span>
<span class="pl-s1">	<span class="pl-k">if</span>(<span class="pl-smi">$pData</span>){</span>
<span class="pl-s1">		<span class="pl-c1">is_array</span>(<span class="pl-smi">$pData</span>) <span class="pl-k">&amp;&amp;</span> <span class="pl-smi">$pData</span> <span class="pl-k">=</span> <span class="pl-c1">http_build_query</span>(<span class="pl-smi">$pData</span>);</span>
<span class="pl-s1">		<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_POST</span>, <span class="pl-c1">true</span>);</span>
<span class="pl-s1">		<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_POSTFIELDS</span>, <span class="pl-smi">$pData</span>);</span>
<span class="pl-s1">	}</span>
<span class="pl-s1">	<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_HTTPHEADER</span>, <span class="pl-c1">array</span>(<span class="pl-s"><span class="pl-pds">&#34;</span>Content-type: application/x-www-form-urlencoded<span class="pl-pds">&#34;</span></span>));</span>
<span class="pl-s1">	<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_URL</span>, <span class="pl-smi">$pUrl</span>);</span>
<span class="pl-s1">	<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_RETURNTRANSFER</span>, <span class="pl-c1">true</span>);</span>
<span class="pl-s1">	<span class="pl-c1">curl_setopt</span>(<span class="pl-smi">$tCh</span>, <span class="pl-c1">CURLOPT_SSL_VERIFYPEER</span>, <span class="pl-c1">false</span>);</span>
<span class="pl-s1">	<span class="pl-smi">$tResult</span> <span class="pl-k">=</span> <span class="pl-c1">curl_exec</span>(<span class="pl-smi">$tCh</span>);</span>
<span class="pl-s1">	<span class="pl-c1">curl_close</span>(<span class="pl-smi">$tCh</span>);</span>
<span class="pl-s1">	<span class="pl-smi">$tmp</span> <span class="pl-k">=</span> <span class="pl-c1">json_decode</span>(<span class="pl-smi">$tResult</span>, <span class="pl-c1">1</span>);</span>
<span class="pl-s1">	<span class="pl-k">if</span>(<span class="pl-smi">$tmp</span>) {</span>
<span class="pl-s1">		<span class="pl-smi">$tResult</span> <span class="pl-k">=</span> <span class="pl-smi">$tmp</span>;</span>
<span class="pl-s1">	}</span>
<span class="pl-s1">	<span class="pl-k">return</span> <span class="pl-smi">$tResult</span>;</span>
<span class="pl-s1">}</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-k">function</span> <span class="pl-en">createSign</span>(<span class="pl-smi">$pParams</span> <span class="pl-k">=</span> <span class="pl-c1">array</span>()){</span>
<span class="pl-s1">	<span class="pl-smi">$pParams</span>[<span class="pl-s"><span class="pl-pds">&#39;</span>secret_key<span class="pl-pds">&#39;</span></span>] <span class="pl-k">=</span> <span class="pl-c1">SECRET_KEY</span>;</span>
<span class="pl-s1">	<span class="pl-c1">ksort</span>(<span class="pl-smi">$pParams</span>);</span>
<span class="pl-s1">	<span class="pl-smi">$tPreSign</span> <span class="pl-k">=</span> <span class="pl-c1">http_build_query</span>(<span class="pl-smi">$pParams</span>);</span>
<span class="pl-s1">	<span class="pl-smi">$tSign</span> <span class="pl-k">=</span> <span class="pl-c1">md5</span>(<span class="pl-smi">$tPreSign</span>);</span>
<span class="pl-s1">	<span class="pl-k">return</span> <span class="pl-c1">strtolower</span>(<span class="pl-smi">$tSign</span>);</span>
<span class="pl-s1">}</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-k">function</span> <span class="pl-en">send2api</span>(<span class="pl-smi">$pParams</span>, <span class="pl-smi">$extra</span> <span class="pl-k">=</span> <span class="pl-c1">array</span>()) {</span>
<span class="pl-s1">	<span class="pl-smi">$pParams</span>[<span class="pl-s"><span class="pl-pds">&#39;</span>access_key<span class="pl-pds">&#39;</span></span>] <span class="pl-k">=</span> <span class="pl-c1">ACCESS_KEY</span>;</span>
<span class="pl-s1">	<span class="pl-smi">$pParams</span>[<span class="pl-s"><span class="pl-pds">&#39;</span>created<span class="pl-pds">&#39;</span></span>] <span class="pl-k">=</span> <span class="pl-c1">time</span>();</span>
<span class="pl-s1">	<span class="pl-smi">$pParams</span>[<span class="pl-s"><span class="pl-pds">&#39;</span>sign<span class="pl-pds">&#39;</span></span>] <span class="pl-k">=</span> createSign(<span class="pl-smi">$pParams</span>);</span>
<span class="pl-s1">	<span class="pl-k">if</span>(<span class="pl-smi">$extra</span>) {</span>
<span class="pl-s1">		<span class="pl-smi">$pParams</span> <span class="pl-k">=</span> <span class="pl-c1">array_merge</span>(<span class="pl-smi">$pParams</span>, <span class="pl-smi">$extra</span>);</span>
<span class="pl-s1">	}</span>
<span class="pl-s1">	<span class="pl-smi">$tResult</span> <span class="pl-k">=</span> httpRequest(<span class="pl-c1">API_URL</span>, <span class="pl-smi">$pParams</span>);</span>
<span class="pl-s1">	<span class="pl-k">return</span> <span class="pl-smi">$tResult</span>;</span>
<span class="pl-s1">}</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-k">function</span> <span class="pl-en">getAccountInfo</span>(){</span>
<span class="pl-s1">	<span class="pl-smi">$tParams</span> <span class="pl-k">=</span> <span class="pl-smi">$extra</span> <span class="pl-k">=</span> <span class="pl-c1">array</span>();</span>
<span class="pl-s1">	<span class="pl-smi">$tParams</span>[<span class="pl-s"><span class="pl-pds">&#39;</span>method<span class="pl-pds">&#39;</span></span>] <span class="pl-k">=</span> <span class="pl-s"><span class="pl-pds">&#39;</span>get_account_info<span class="pl-pds">&#39;</span></span>;</span>
<span class="pl-s1">	<span class="pl-c"><span class="pl-c">//</span> 不参与签名样例</span></span>
<span class="pl-s1">	<span class="pl-c"><span class="pl-c">//</span> $extra[&#39;test&#39;] = &#39;test&#39;;</span></span>
<span class="pl-s1">	<span class="pl-smi">$tResult</span> <span class="pl-k">=</span> send2api(<span class="pl-smi">$tParams</span>, <span class="pl-smi">$extra</span>);</span>
<span class="pl-s1">	<span class="pl-k">return</span> <span class="pl-smi">$tResult</span>;</span>
<span class="pl-s1">}</span>
<span class="pl-s1"></span>
<span class="pl-s1"><span class="pl-k">try</span> {</span>
<span class="pl-s1">	<span class="pl-c1">var_dump</span>(getAccountInfo());</span>
<span class="pl-s1">} <span class="pl-k">catch</span> (<span class="pl-c1">Exception</span> <span class="pl-smi">$e</span>) {</span>
<span class="pl-s1">	<span class="pl-c1">var_dump</span>(<span class="pl-smi">$e</span>);</span>
<span class="pl-s1">}</span></pre></div>

        