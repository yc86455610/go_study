
          <p>##分时行情数据接口（K线）
<code>目前支持人民币现货、美元现货</code>
###数据文件：
[BTC-CNY] <code>http://api.huobi.com/staticmarket/btc_kline_[period]_json.js</code></p>
<p>[LTC-CNY] <code>http://api.huobi.com/staticmarket/ltc_kline_[period]_json.js</code></p>
<p>[BTC-USD] <code>http://api.huobi.com/usdmarket/btc_kline_[period]_json.js</code></p>
<table>
    <thead>
    <tr>
        <th width="25%">period 参数</th>
        <th>说明</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <th>001</th>
        <td>1分钟线</td>
    </tr>
    <tr>
        <th>005</th>
        <td>5分钟</td>
    </tr>
    <tr>
        <th>015</th>
        <td>15分钟</td>
    </tr>
    <tr>
        <th>030</th>
        <td>30分钟</td>
    </tr>
    <tr>
        <th>060</th>
        <td>60分钟</td>
    </tr>
    <tr>
        <th>100</th>
        <td>日线</td>
    </tr>
    <tr>
        <th>200</th>
        <td>周线</td>
    </tr>
    <tr>
        <th>300</th>
        <td>月线</td>
    </tr>
    <tr>
        <th>400</th>
        <td>年线</td>
    </tr>
    <tr>
    	<th>例如</th>
        <td>
            [BTC-CNY] http://api.huobi.com/staticmarket/btc_kline_005_json.js<br/>
            [LTC-LTC] http://api.huobi.com/staticmarket/ltc_kline_005_json.js<br/>
            [BTC-USD] http://api.huobi.com/usdmarket/btc_kline_005_json.js
        </td>
    </tr>
    </tbody>
</table>
<table>
    <thead>
    <tr>
        <th width="25%">length参数</th>
        <th>说明</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <th>1~2000</th>
        <td>返回1~2000条数据，如果不传该参数默认为300条</td>
    </tr>
    <tr>
    	<th>例如</th>
        <td>
            [BTC-CNY] http://api.huobi.com/staticmarket/btc_kline_005_json.js?length=500<br/>
<pre><code>    &lt;/td&gt;
&lt;/tr&gt;
&lt;/tbody&gt;
</code></pre>
</td>
</tr>
</tbody>
</table>
<p>###数据格式:（json）</p>
<div class="highlight highlight-source-js"><pre>[[<span class="pl-s"><span class="pl-pds">&#34;</span>20140417095000000<span class="pl-pds">&#34;</span></span>,<span class="pl-c1">3297</span>,<span class="pl-c1">3305</span>,<span class="pl-c1">3290</span>,<span class="pl-c1">3303</span>,<span class="pl-c1">386.3926</span>],<span class="pl-k">...</span>]</pre></div>
<p>日期时间，开盘价，最高价，最低价，收盘价，成交量
###用例：
<img src="https://camo.githubusercontent.com/c894614f72cb0d2c32cd8e171b69a98a5f3d0e15/68747470733a2f2f7374617469632e68756f62692e636f6d2f696d672f68656c702f6d61726b65745f68656c705f696d67312e706e67" alt="行情图例" data-canonical-src="https://static.huobi.com/img/help/market_help_img1.png"/></p>

        