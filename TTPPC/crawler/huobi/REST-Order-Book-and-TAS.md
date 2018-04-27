
          <p>##买卖盘实时成交数据
<code>目前支持人民币现货、美元现货</code>
###数据文件：
[BTC-CNY] <code>http://api.huobi.com/staticmarket/detail_btc_json.js</code></p>
<p>[LTC-CNY] <code>http://api.huobi.com/staticmarket/detail_ltc_json.js</code></p>
<p>[BTC-USD] <code>http://api.huobi.com/usdmarket/detail_btc_json.js</code></p>
<p>###数据格式:(json)：</p>
<div class="highlight highlight-source-js"><pre>{
  amount<span class="pl-k">:</span> <span class="pl-c1">63165</span> <span class="pl-c"><span class="pl-c">//</span>成交量</span>
  level<span class="pl-k">:</span> <span class="pl-c1">86.999</span> <span class="pl-c"><span class="pl-c">//</span>涨幅</span>
  buys<span class="pl-k">:</span> <span class="pl-c1">Array</span>[<span class="pl-c1">10</span>] <span class="pl-c"><span class="pl-c">//</span>买10</span>
  p_high<span class="pl-k">:</span> <span class="pl-c1">4410</span> <span class="pl-c"><span class="pl-c">//</span>最高</span>
  p_last<span class="pl-k">:</span> <span class="pl-c1">4275</span> <span class="pl-c"><span class="pl-c">//</span>收盘价</span>
  p_low<span class="pl-k">:</span> <span class="pl-c1">4250</span> <span class="pl-c"><span class="pl-c">//</span>最低</span>
  p_new<span class="pl-k">:</span> <span class="pl-c1">4362</span> <span class="pl-c"><span class="pl-c">//</span>最新</span>
  p_open<span class="pl-k">:</span> <span class="pl-c1">4275</span> <span class="pl-c"><span class="pl-c">//</span>开盘</span>
  sells<span class="pl-k">:</span> <span class="pl-c1">Array</span>[<span class="pl-c1">10</span>] <span class="pl-c"><span class="pl-c">//</span>卖10</span>
  top_buy<span class="pl-k">:</span> <span class="pl-c1">Array</span>[<span class="pl-c1">5</span>] <span class="pl-c"><span class="pl-c">//</span>买5</span>
  top_sell<span class="pl-k">:</span> <span class="pl-c1">Array</span>[<span class="pl-c1">5</span>] <span class="pl-c"><span class="pl-c">//</span>卖5 </span>
  total<span class="pl-k">:</span> <span class="pl-c1">273542407.24361</span> <span class="pl-c"><span class="pl-c">//</span>总量（人民币） </span>
  trades<span class="pl-k">:</span> <span class="pl-c1">Array</span>[<span class="pl-c1">15</span>] <span class="pl-c"><span class="pl-c">//</span>实时成交</span>
  symbol<span class="pl-k">:</span><span class="pl-s"><span class="pl-pds">&#34;</span>btccny<span class="pl-pds">&#34;</span></span> <span class="pl-c"><span class="pl-c">//</span>类型</span>
}</pre></div>
<p>买卖盘用例：</p>
<p><img src="https://camo.githubusercontent.com/574d2bcf1c485642ffe83d80e622e972311a49c9/68747470733a2f2f7374617469632e68756f62692e636f6d2f696d672f68656c702f6d61726b65745f68656c705f696d67332e706e67" alt="买卖盘用例 " data-canonical-src="https://static.huobi.com/img/help/market_help_img3.png"/></p>
<p>实时成交用例:</p>
<p><img src="https://camo.githubusercontent.com/1891910fc5513093fbac7021ddedc93a7a99b5f7/68747470733a2f2f7374617469632e68756f62692e636f6d2f696d672f68656c702f6d61726b65745f68656c705f696d67342e706e67" alt="实时成交用例 " data-canonical-src="https://static.huobi.com/img/help/market_help_img4.png"/></p>

        