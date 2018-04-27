
          <p>##实时行情数据接口
<code>目前支持人民币现货、美元现货</code>
###数据文件：
[BTC-CNY] <code>http://api.huobi.com/staticmarket/ticker_btc_json.js</code></p>
<p>[LTC-CNY] <code>http://api.huobi.com/staticmarket/ticker_ltc_json.js</code></p>
<p>[BTC-USD] <code>http://api.huobi.com/usdmarket/ticker_btc_json.js</code></p>
<p>###数据格式:</p>
<div class="highlight highlight-source-js"><pre>{<span class="pl-s"><span class="pl-pds">&#34;</span>time<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-s"><span class="pl-pds">&#34;</span>1378137600<span class="pl-pds">&#34;</span></span>,<span class="pl-s"><span class="pl-pds">&#34;</span>ticker<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span>{<span class="pl-s"><span class="pl-pds">&#34;</span>high<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">86.48</span>,<span class="pl-s"><span class="pl-pds">&#34;</span>low<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">79.75</span>,<span class="pl-s"><span class="pl-pds">&#34;</span>symbol<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-s"><span class="pl-pds">&#34;</span>btccny<span class="pl-pds">&#34;</span></span>,<span class="pl-s"><span class="pl-pds">&#34;</span>last<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">83.9</span>,<span class="pl-s"><span class="pl-pds">&#34;</span>vol<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">2239560.1752883</span>,<span class="pl-s"><span class="pl-pds">&#34;</span>buy<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">83.88</span>,<span class="pl-s"><span class="pl-pds">&#34;</span>sell<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-c1">83.9</span>}}</pre></div>
<p>报价：最高价，最低价，当前价，成交量，买1，卖1
##深度数据接口（json格式）
###数据文件
[BTC-CNY] <code>http://api.huobi.com/staticmarket/depth_btc_json.js</code></p>
<p>[LTC-CNY] <code>http://api.huobi.com/staticmarket/depth_ltc_json.js</code></p>
<p>[BTC-USD] <code>http://api.huobi.com/usdmarket/depth_btc_json.js</code></p>
<p>指定深度数据条数（1-150条）</p>
<p>[BTC-CNY] <code>http://api.huobi.com/staticmarket/depth_btc_X.js</code></p>
<p>[LTC-CNY] <code>http://api.huobi.com/staticmarket/depth_ltc_X.js</code></p>
<p>[BTC-USD] <code>http://api.huobi.com/usdmarket/depth_btc_X.js</code></p>
<p>X表示返回多少条深度数据，可取值 1-150</p>
<p>###数据格式</p>
<div class="highlight highlight-source-js"><pre>{<span class="pl-s"><span class="pl-pds">&#34;</span>asks<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span>[[<span class="pl-c1">90.8</span>,<span class="pl-c1">0.5</span>],<span class="pl-k">...</span>],<span class="pl-s"><span class="pl-pds">&#34;</span>bids<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span>[[<span class="pl-c1">86.06</span>,<span class="pl-c1">79.243</span>],<span class="pl-k">...</span>]],<span class="pl-s"><span class="pl-pds">&#34;</span>symbol<span class="pl-pds">&#34;</span></span><span class="pl-k">:</span><span class="pl-s"><span class="pl-pds">&#34;</span>btccny<span class="pl-pds">&#34;</span></span>}</pre></div>
<p>卖:价格:累积量,...    买:价格:累积量...</p>

        