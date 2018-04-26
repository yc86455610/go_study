<h1><a id="user-content-rest-api-for-futures--" class="anchor" aria-hidden="true" href="#rest-api-for-futures--"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>REST API for FUTURES    </h1>
<h2><a id="user-content-开始使用--" class="anchor" aria-hidden="true" href="#开始使用--"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>开始使用    </h2>
<p>REST，即Representational State Transfer的缩写，是目前最流行的一种互联网软件架构。它结构清晰、符合标准、易于理解、扩展方便，正得到越来越多网站的采用。其优点如下：</p>
<ul>
<li>在RESTful架构中，每一个URL代表一种资源；</li>
<li>客户端和服务器之间，传递这种资源的某种表现层；</li>
<li>客户端通过四个HTTP指令，对服务器端资源进行操作，实现“表现层状态转化”。</li>
</ul>
<p>建议开发者使用REST API进行币币交易或者资产提现等操作。</p>
<h2><a id="user-content-请求交互--" class="anchor" aria-hidden="true" href="#请求交互--"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>请求交互    </h2>
<p>REST访问的根URL：<code>https://www.okex.com/api/v1</code></p>
<p>所有请求基于Https协议，请求头信息中contentType需要统一设置为：<code>application/x-www-form-urlencoded</code></p>
<p>请求交互说明</p>
<ol>
<li>请求参数：根据接口请求参数规定进行参数封装。</li>
<li>提交请求参数：将封装好的请求参数通过POST 或GET 方式提交至服务器。</li>
<li>服务器响应：服务器首先对用户请求数据进行参数安全校验，通过校验后根据业务逻辑将响应数据以JSON格式返回给用户。</li>
<li>数据处理：对服务器响应数据进行处理。</li>
</ol>
<h2><a id="user-content-api参考" class="anchor" aria-hidden="true" href="#api参考"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>API参考</h2>
<h3><a id="user-content-合约行情-api" class="anchor" aria-hidden="true" href="#合约行情-api"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>合约行情 API</h3>
<p>获取OKEx合约行情数据</p>
<ol>
<li>Get /api/v1/future_ticker    获取OKEx合约行情</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_ticker.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_ticker.do?symbol=btc_usd&amp;contract_type=this_week
# Response
{
	&#34;date&#34;:&#34;1411627632&#34;,
	&#34;ticker&#34;:{
		&#34;last&#34;:409.2, 
		&#34;buy&#34; :408.23, 
		&#34;sell&#34;:409.18, 
		&#34;high&#34;:432.0, 
		&#34;low&#34;:405.71, 
		&#34;vol&#34;:55168.0, 
		&#34;contract_id&#34;:20140926012, 
		&#34;unit_amount&#34;:100.0
	}
}
</code></pre>
<p>返回值说明</p>
<pre><code>buy:买一价
contract_id:合约ID
high:最高价
last:最新成交价
low:最低价
sell:卖一价
unit_amount:合约面值
vol:成交量(最近的24小时)
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr></tbody></table>
<ol start="2">
<li>Get /api/v1/future_depth   获取OKEx合约深度信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_depth.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_depth.do?symbol=btc_usd&amp;contract_type=this_week
# Response
{
	&#34;asks&#34;:[
		[411.8,6],
		[411.75,11],
		[411.6,22],
		[411.5,9],
		[411.3,16]
	],
	&#34;bids&#34;:[
		[410.65,12],
		[410.64,3],
		[410.19,15],
		[410.18,40],
		[410.09,10]
	]
}
</code></pre>
<p>返回值说明</p>
<pre><code>asks :卖方深度
bids :买方深度
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">size</td>
<td align="left">Integer</td>
<td align="left">是</td>
<td align="left">value：1-200</td>
</tr>
<tr>
<td align="left">merge</td>
<td align="left">Integer</td>
<td align="left">否(默认0)</td>
<td align="left">value：1(合并深度)</td>
</tr></tbody></table>
<ol start="3">
<li>Get /api/v1/future_trades   获取OKEx合约交易记录信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_trades.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_trades.do?symbol=btc_usd&amp;contract_type=this_week
# Response
[
	{
		&#34;amount&#34;:11,
		&#34;date_ms&#34;:140807646000,
		&#34;date&#34;:140807646,
		&#34;price&#34;:7.076,
		&#34;tid&#34;:37,
		&#34;type&#34;:&#34;buy&#34;
	},
	{
		&#34;amount&#34;:100,
		&#34;date_ms&#34;:1408076464000,
		&#34;date&#34;:1408076464,
		&#34;price&#34;:7.076,
		&#34;tid&#34;:39,
		&#34;type&#34;:&#34;sell&#34;
	}
]
</code></pre>
<p>返回值说明</p>
<pre><code>amount：交易数量
date_ms：交易时间(毫秒)
date：交易时间
price：交易价格
tid：交易ID
type：交易类型
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr></tbody></table>
<ol start="4">
<li>Get /api/v1/future_index   获取OKEx合约指数信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_index.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_index.do?symbol=btc_usd
# Response
{&#34;future_index&#34;:471.0817}
</code></pre>
<p>返回值说明</p>
<pre><code>future_index :指数
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr></tbody></table>
<ol start="5">
<li>Get /api/v1/exchange_rate   获取美元人民币汇率</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/exchange_rate.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/exchange_rate.do
# Response
{ &#34;rate&#34;:6.1329 }
</code></pre>
<p>返回值说明</p>
<pre><code>rate：美元-人民币汇率
</code></pre>
<p>请求参数</p>
<pre><code>无
</code></pre>
<ol start="6">
<li>Get /api/v1/future_estimated_price   获取交割预估价</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_estimated_price.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_depth.do?symbol=btc_usd
# Response
{&#34;forecast_price&#34;:5.4}
</code></pre>
<p>返回值说明</p>
<pre><code>forecast_price:交割预估价  注意：交割预估价只有交割前三小时返回
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr></tbody></table>
<ol start="7">
<li>Get /api/v1/future_kline   获取OKEx合约深度信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_kline.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_kline.do
# Response
[
    [
        1440308700000,
        233.37,
        233.48,
        233.37,
        233.48,
        52,
        22.2810015
    ],
    [
        1440308760000,
        233.38,
        233.38,
        233.27,
        233.37,
        186,
        79.70234956
    ]
]
</code></pre>
<p>返回值说明</p>
<pre><code>[
    1440308760000,	时间戳
    233.38,		开
    233.38,		高
    233.27,		低
    233.37,		收
    186,		交易量
    79.70234956		交易量转化BTC或LTC数量
]
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">1min/3min/5min/15min/30min/1day/3day/1week/1hour/2hour/4hour/6hour/12hour</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">size</td>
<td align="left">Integer</td>
<td align="left">否(默认0)</td>
<td align="left">指定获取数据的条数</td>
</tr>
<tr>
<td align="left">since</td>
<td align="left">Long</td>
<td align="left">否(默认0)</td>
<td align="left">时间戳（eg：1417536000000）。 返回该时间戳以后的数据</td>
</tr></tbody></table>
<ol start="8">
<li>Get /api/v1/future_hold_amount   获取当前可用合约总持仓量</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_hold_amount.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_hold_amount.do?symbol=btc_usd&amp;contract_type=this_week
# Response
[
    {
        &#34;amount&#34;: 106856,
        &#34;contract_name&#34;: &#34;BTC0213&#34;
    }
]
</code></pre>
<p>返回值说明</p>
<pre><code>amount:总持仓量（张）
contract_name:合约名
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr></tbody></table>
<ol start="9">
<li>Get /api/v1/future_price_limit   获取合约最高限价和最低限价</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_price_limit.do</code></p>
<p>示例</p>
<pre><code># Request
GET https://www.okex.com/api/v1/future_price_limit.do?symbol=btc_usd&amp;contract_type=this_week
# Response
{&#34;high&#34;:443.07,&#34;low&#34;:417.09}
</code></pre>
<p>返回值说明</p>
<pre><code>high :最高买价
low :最低卖价
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr></tbody></table>
<h3><a id="user-content-合约交易-api" class="anchor" aria-hidden="true" href="#合约交易-api"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>合约交易 API</h3>
<p>用于OKEX快速进行合约交易</p>
<ol>
<li>POST /api/v1/future_userinfo   获取OKEx合约账户信息(全仓)</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_userinfo.do</code>  访问频率 10次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_userinfo.do
# Response
{
    &#34;info&#34;: {
        &#34;btc&#34;: {
            &#34;account_rights&#34;: 1,
            &#34;keep_deposit&#34;: 0,
            &#34;profit_real&#34;: 3.33,
            &#34;profit_unreal&#34;: 0,
            &#34;risk_rate&#34;: 10000
        },
        &#34;ltc&#34;: {
            &#34;account_rights&#34;: 2,
            &#34;keep_deposit&#34;: 2.22,
            &#34;profit_real&#34;: 3.33,
            &#34;profit_unreal&#34;: 2,
            &#34;risk_rate&#34;: 10000
        }
    },
    &#34;result&#34;: true
}
</code></pre>
<p>返回值说明</p>
<pre><code>account_rights:账户权益
keep_deposit：保证金
profit_real：已实现盈亏
profit_unreal：未实现盈亏
risk_rate：保证金率

</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr></tbody></table>
<ol start="2">
<li>POST /api/v1/future_position   获取用户持仓获取OKEX合约账户信息 （全仓）</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_position.do</code>  访问频率 10次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_position.do
# Response
{
	&#34;force_liqu_price&#34;: &#34;0.07&#34;,
    &#34;holding&#34;: [
        {
            &#34;buy_amount&#34;: 1,
            &#34;buy_available&#34;: 0,
            &#34;buy_price_avg&#34;: 422.78,
            &#34;buy_price_cost&#34;: 422.78,
            &#34;buy_profit_real&#34;: -0.00007096,
            &#34;contract_id&#34;: 20141219012,
            &#34;contract_type&#34;: &#34;this_week&#34;,
            &#34;create_date&#34;: 1418113356000,
            &#34;lever_rate&#34;: 10,
            &#34;sell_amount&#34;: 0,
            &#34;sell_available&#34;: 0,
            &#34;sell_price_avg&#34;: 0,
            &#34;sell_price_cost&#34;: 0,
            &#34;sell_profit_real&#34;: 0,
            &#34;symbol&#34;: &#34;btc_usd&#34;
        }
    ],
    &#34;result&#34;: true
}
</code></pre>
<p>返回值说明</p>
<pre><code>buy_amount(double):多仓数量
buy_available:多仓可平仓数量
buy_price_avg(double):开仓平均价
buy_price_cost(double):结算基准价
buy_profit_real(double):多仓已实现盈余
contract_id(long):合约id
create_date(long):创建日期
lever_rate:杠杆倍数
sell_amount(double):空仓数量
sell_available:空仓可平仓数量
sell_price_avg(double):开仓平均价
sell_price_cost(double):结算基准价
sell_profit_real(double):空仓已实现盈余
symbol:btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd
contract_type:合约类型
force_liqu_price:预估爆仓价
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr></tbody></table>
<ol start="3">
<li>POST /api/v1/future_trade   合约下单</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_trade.do</code>  访问频率 20次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_trade.do
# Response
{
	&#34;order_id&#34;:986,
   &#34;result&#34;:true
}
</code></pre>
<p>返回值说明</p>
<pre><code>order_id ： 订单ID
result ： true代表成功返回
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">price</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">价格</td>
</tr>
<tr>
<td align="left">amount</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">委托数量</td>
</tr>
<tr>
<td align="left">type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">1:开多 2:开空 3:平多 4:平空</td>
</tr>
<tr>
<td align="left">match_price</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">是否为对手价 0:不是    1:是   ,当取值为1时,price无效</td>
</tr>
<tr>
<td align="left">lever_rate</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">杠杆倍数，下单时无需传送，系统取用户在页面上设置的杠杆倍数。且“开仓”若有10倍多单，就不能再下20倍多单</td>
</tr></tbody></table>
<ol start="4">
<li>POST /api/v1/future_trades_history    获取OKEX合约交易历史（非个人）访问频率</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_trades_history</code>   访问频率 2次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_trades_history.do
# Response
[
    {
        &#34;amount&#34;: 11,
        &#34;date&#34;: 140807646000,
        &#34;price&#34;: 7.076,
        &#34;tid&#34;: 37,
        &#34;type&#34;: &#34;buy&#34;
    },
    {
        &#34;amount&#34;: 100,
        &#34;date&#34;: 1408076464000,
        &#34;price&#34;: 7.076,
        &#34;tid&#34;: 39,
        &#34;type&#34;: &#34;sell&#34;
    }
]
</code></pre>
<p>返回值说明</p>
<pre><code>amount：交易数量
date：交易时间(毫秒)
price：交易价格
tid：交易ID
type：交易类型（buy/sell）
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">date</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约交割时间，格式yyyy-MM-dd</td>
</tr>
<tr>
<td align="left">since</td>
<td align="left">Long</td>
<td align="left">是</td>
<td align="left">交易Id起始位置</td>
</tr></tbody></table>
<ol start="5">
<li>POST /api/v1/future_batch_trade   批量下单</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_batch_trade.do</code>  访问频率 10次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_batch_trade.do
# Response
{
	&#34;order_info&#34;:[
        {&#34;error_code&#34;:20013,&#34;order_id&#34;:-1},
        {&#34;error_code&#34;:20013,&#34;order_id&#34;:-1},
        {&#34;order_id&#34;:161256},
        {&#34;order_id&#34;:161257}
	],
	&#34;result&#34;:true
}
</code></pre>
<p>返回值说明</p>
<pre><code>result:订单交易成功或失败
order_id:订单ID
备注:
     只要其中任何一单下单成功就返回true
      返回的结果数据和orders_data提交订单数据顺序一致
     如果下单失败：order_id为-1，error_code为错误代码
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">orders_data</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">JSON类型的字符串 例：[{price:5,amount:2,type:1,match_price:1},{price:2,amount:3,type:1,match_price:1}] 最大下单量为5，price,amount,type,match_price参数参考future_trade接口中的说明</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">lever_rate</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">杠杆倍数，下单时无需传送，系统取用户在页面上设置的杠杆倍数。且“开仓”若有10倍多单，就不能再下20倍多单</td>
</tr></tbody></table>
<ol start="6">
<li>POST /api/v1/future_cancel   取消合约订单</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_cancel.do</code>  访问频率 4次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_cancel.do
# Response
#多笔订单返回结果(成功订单ID,失败订单ID)
{
    &#34;error&#34;:&#34;161251:20015&#34;, //161251订单id 20015错误id
    &#34;success&#34;:&#34;161256&#34;
}
#单笔订单返回结果
{
    &#34;order_id&#34;:&#34;161277&#34;,
    &#34;result&#34;:true
}
</code></pre>
<p>返回值说明</p>
<pre><code>result:订单交易成功或失败(用于单笔订单)
order_id:订单ID(用于单笔订单)
success:成功的订单ID(用于多笔订单)
error:失败的订单ID后跟失败错误码(用户多笔订单)
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">order_id</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">订单ID(多个订单ID中间以&#34;,&#34;分隔,一次最多允许撤消3个订单)</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr></tbody></table>
<ol start="7">
<li>POST /api/v1/future_order_info   获取合约订单信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_order_info.do</code></p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_order_info.do
# Response
{
  &#34;orders&#34;:
     [
        {
            &#34;amount&#34;:111,
            &#34;contract_name&#34;:&#34;LTC0815&#34;,
            &#34;create_date&#34;:1408076414000,
            &#34;deal_amount&#34;:1,
            &#34;fee&#34;:0,
            &#34;order_id&#34;:106837,
            &#34;price&#34;:1111,
            &#34;price_avg&#34;:0,
            &#34;status&#34;:&#34;0&#34;,
            &#34;symbol&#34;:&#34;ltc_usd&#34;,
            &#34;type&#34;:&#34;1&#34;,
            &#34;unit_amount&#34;:100,
            &#34;lever_rate&#34;:10
        }
     ],
   &#34;result&#34;:true
}
</code></pre>
<p>返回值说明</p>
<pre><code>amount: 委托数量
contract_name: 合约名称
create_date: 委托时间
deal_amount: 成交数量
fee: 手续费
order_id: 订单ID
price: 订单价格
price_avg: 平均价格
status: 订单状态(0等待成交 1部分成交 2全部成交 -1撤单 4撤单处理中 5撤单中)
symbol: btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd
type: 订单类型 1：开多 2：开空 3：平多 4： 平空
unit_amount:合约面值
lever_rate: 杠杆倍数  value:10\20  默认10 
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">status</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">查询状态 1:未完成的订单 2:已经完成的订单</td>
</tr>
<tr>
<td align="left">order_id</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">订单ID -1:查询指定状态的订单，否则查询相应订单号的订单</td>
</tr>
<tr>
<td align="left">current_page</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">当前页数</td>
</tr>
<tr>
<td align="left">page_length</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">每页获取条数，最多不超过50</td>
</tr></tbody></table>
<ol start="8">
<li>POST /api/v1/future_orders_info   批量获取合约订单信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_orders_info.do</code></p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_orders_info.do
# Response
{
    &#34;orders&#34;: [
        {
            &#34;amount&#34;: 1,
            &#34;contract_name&#34;: &#34;BTC0213&#34;,
            &#34;create_date&#34;: 1424932853000,
            &#34;deal_amount&#34;: 0,
            &#34;fee&#34;: 0,
            &#34;lever_rate&#34;: 20,
            &#34;order_id&#34;: 200144,
            &#34;price&#34;: 1,
            &#34;price_avg&#34;: 0,
            &#34;status&#34;: 0,
            &#34;symbol&#34;: &#34;btc_usd&#34;,
            &#34;type&#34;: 1,
            &#34;unit_amount&#34;: 100
        },
        {
            &#34;amount&#34;: 1,
            &#34;contract_name&#34;: &#34;BTC0213&#34;,
            &#34;create_date&#34;: 1424932903000,
            &#34;deal_amount&#34;: 0,
            &#34;fee&#34;: 0,
            &#34;lever_rate&#34;: 20,
            &#34;order_id&#34;: 200146,
            &#34;price&#34;: 1,
            &#34;price_avg&#34;: 0,
            &#34;status&#34;: 0,
            &#34;symbol&#34;: &#34;btc_usd&#34;,
            &#34;type&#34;: 1,
            &#34;unit_amount&#34;: 100
        },
        {
            &#34;amount&#34;: 1,
            &#34;contract_name&#34;: &#34;BTC0213&#34;,
            &#34;create_date&#34;: 1424932917000,
            &#34;deal_amount&#34;: 0,
            &#34;fee&#34;: 0,
            &#34;lever_rate&#34;: 20,
            &#34;order_id&#34;: 200147,
            &#34;price&#34;: 1,
            &#34;price_avg&#34;: 0,
            &#34;status&#34;: 0,
            &#34;symbol&#34;: &#34;btc_usd&#34;,
            &#34;type&#34;: 1,
            &#34;unit_amount&#34;: 100
        }
    ],
    &#34;result&#34;: true
}
</code></pre>
<p>返回值说明</p>
<pre><code>amount: 委托数量
contract_name: 合约名称
created_date: 委托时间
deal_amount: 成交数量
fee: 手续费
order_id: 订单ID
price: 订单价格
price_avg: 平均价格
status: 订单状态(0等待成交 1部分成交 2全部成交 -1撤单 4撤单处理中)
symbol: btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd
type: 订单类型 1：开多 2：开空 3：平多 4： 平空
unit_amount:合约面值
lever_rate: 杠杆倍数  value:10\20  默认10  
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">order_id</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">订单ID(多个订单ID中间以&#34;,&#34;分隔,一次最多允许查询50个订单)</td>
</tr></tbody></table>
<ol start="9">
<li>POST /api/v1/future_userinfo_4fix   获取逐仓合约账户信息</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_userinfo_4fix.do</code></p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_userinfo_4fix.do
# Response
{
	&#34;info&#34;: {
        &#34;btc&#34;: {
            &#34;balance&#34;: 99.95468925,
            &#34;contracts&#34;: [{
                &#34;available&#34;: 99.95468925,
                &#34;balance&#34;: 0.03779061,
                &#34;bond&#34;: 0,
                &#34;contract_id&#34;: 20140815012,
                &#34;contratct_type&#34;: &#34;this_week&#34;,
                &#34;freeze&#34;: 0,
                &#34;profit&#34;: -0.03779061,
                &#34;unprofit&#34;: 0
            }],
            &#34;rights&#34;: 99.95468925
        },
        &#34;ltc&#34;: {
            &#34;balance&#34;: 77,
            &#34;contracts&#34;: [{
                &#34;available&#34;: 99.95468925,
                &#34;balance&#34;: 0.03779061,
                &#34;bond&#34;: 0,
                &#34;contract_id&#34;: 20140815012,
                &#34;contract_type&#34;: &#34;this_week&#34;,
                &#34;freeze&#34;: 0,
                &#34;profit&#34;: -0.03779061,
                &#34;unprofit&#34;: 0
            }],
            &#34;rights&#34;: 100
        }
    },
    &#34;result&#34;: true
}
</code></pre>
<p>返回值说明</p>
<pre><code>balance:账户余额
available:合约可用
balance:账户(合约)余额
bond:固定保证金
contract_id:合约ID
contract_type:合约类别
freeze:冻结
profit:已实现盈亏
unprofit:未实现盈亏
rights:账户权益
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr></tbody></table>
<ol start="10">
<li>POST /api/v1/future_position_4fix   逐仓用户持仓查询</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_position_4fix.do</code>  访问频率 10次/2秒</p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_position_4fix.do
# Response
{
	&#34;holding&#34;: [{
        &#34;buy_amount&#34;: 10,
        &#34;buy_available&#34;: 2,
        &#34;buy_bond&#34;: 1.27832803,
        &#34;buy_flatprice&#34;: &#34;338.97&#34;,
        &#34;buy_price_avg&#34;: 555.67966869,
        &#34;buy_price_cost&#34;: 555.67966869,
        &#34;buy_profit_lossratio&#34;: &#34;13.52&#34;,
        &#34;buy_profit_real&#34;: 0,
        &#34;contract_id&#34;: 20140815012,
        &#34;contract_type&#34;: &#34;this_week&#34;,
        &#34;create_date&#34;: 1408594176000,
        &#34;sell_amount&#34;: 8,
        &#34;sell_available&#34;: 2,
        &#34;sell_bond&#34;: 0.24315591,
        &#34;sell_flatprice&#34;: &#34;671.15&#34;,
        &#34;sell_price_avg&#34;: 567.04644056,
        &#34;sell_price_cost&#34;: 567.04644056,
        &#34;sell_profit_lossratio&#34;: &#34;-45.04&#34;,
        &#34;sell_profit_real&#34;: 0,
        &#34;symbol&#34;: &#34;btc_usd&#34;,
        &#34;lever_rate&#34;: 10
    }],
    &#34;result&#34;: true
}
</code></pre>
<p>返回值说明</p>
<pre><code>buy_amount:多仓数量
buy_available:多仓可平仓数量 
buy_bond:多仓保证金
buy_flatprice:多仓强平价格
buy_profit_lossratio:多仓盈亏比
buy_price_avg:开仓平均价
buy_price_cost:结算基准价
buy_profit_real:多仓已实现盈余
contract_id:合约id
contract_type:合约类型
create_date:创建日期
sell_amount:空仓数量
sell_available:空仓可平仓数量 
sell_bond:空仓保证金
sell_flatprice:空仓强平价格
sell_profit_lossratio:空仓盈亏比
sell_price_avg:开仓平均价
sell_price_cost:结算基准价
sell_profit_real:空仓已实现盈余
symbol:btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd
lever_rate: 杠杆倍数
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">type</td>
<td align="left">String</td>
<td align="left">否</td>
<td align="left">默认返回10倍杠杆持仓 type=1 返回全部持仓数据</td>
</tr></tbody></table>
<ol start="11">
<li>POST /api/v1/future_explosive   获取合约爆仓单</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_explosive.do</code></p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_explosive.do
# Response
[
    {
        &#34;data&#34;: [
            {
                &#34;amount&#34;: &#34;42&#34;,
                &#34;create_date&#34;: &#34;2015-02-27 11:48:07&#34;,
                &#34;loss&#34;: &#34;-0.54275722&#34;,
                &#34;price&#34;: &#34;249.02&#34;,
                &#34;type&#34;: 4
            },
            {
                &#34;amount&#34;: &#34;12&#34;,
                &#34;create_date&#34;: &#34;2015-02-27 11:48:05&#34;,
                &#34;loss&#34;: &#34;-0.15507349&#34;,
                &#34;price&#34;: &#34;249.02&#34;,
                &#34;type&#34;: 4
            },
            {
                &#34;amount&#34;: &#34;3&#34;,
                &#34;create_date&#34;: &#34;2015-02-27 11:48:01&#34;,
                &#34;loss&#34;: &#34;-0.03896192&#34;,
                &#34;price&#34;: &#34;248.98&#34;,
                &#34;type&#34;: 4
            },
            {
                &#34;amount&#34;: &#34;1811&#34;,
                &#34;create_date&#34;: &#34;2015-02-27 11:48:01&#34;,
                &#34;loss&#34;: &#34;-23.40317457&#34;,
                &#34;price&#34;: &#34;249.02&#34;,
                &#34;type&#34;: 4
            },
            {
                &#34;amount&#34;: &#34;2&#34;,
                &#34;create_date&#34;: &#34;2015-02-27 11:48:00&#34;,
                &#34;loss&#34;: &#34;-0.02655576&#34;,
                &#34;price&#34;: &#34;248.80&#34;,
                &#34;type&#34;: 4
            }
        ]
    }
]
</code></pre>
<p>返回值说明</p>
<pre><code>amount:委托数量（张）
create_date:创建时间
loss:穿仓用户亏损
price:委托价格
type：交易类型 1：买入开多 2：卖出开空 3：卖出平多 4：买入平空
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">contract_type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">合约类型: this_week:当周   next_week:下周   quarter:季度</td>
</tr>
<tr>
<td align="left">status</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">状态 0：最近7天未成交 1:最近7天已成交</td>
</tr>
<tr>
<td align="left">current_page</td>
<td align="left">Integer</td>
<td align="left">否</td>
<td align="left">当前页数索引值</td>
</tr>
<tr>
<td align="left">page_number</td>
<td align="left">Integer</td>
<td align="left">否</td>
<td align="left">当前页数(使用page_number时current_page失效，current_page无需传)</td>
</tr>
<tr>
<td align="left">page_length</td>
<td align="left">Integer</td>
<td align="left">否</td>
<td align="left">每页获取条数，最多不超过50</td>
</tr></tbody></table>
<ol start="12">
<li>POST /api/v1/future_devolve   个人账户资金划转</li>
</ol>
<p>URL <code>https://www.okex.com/api/v1/future_devolve.do</code></p>
<p>示例</p>
<pre><code># Request
POST https://www.okex.com/api/v1/future_devolve.do
# Response
{
    &#34;result&#34;:true
}
或
{
    &#34;error_code&#34;:20029,
    &#34;result&#34;:false
}
</code></pre>
<p>返回值说明</p>
<pre><code>result:划转结果。若是划转失败，将给出错误码提示。
</code></pre>
<p>请求参数</p>
<table>
<thead>
<tr>
<th align="left">参数名</th>
<th align="left">参数类型</th>
<th align="left">必填</th>
<th align="left">描述</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">api_key</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">用户申请的apiKey</td>
</tr>
<tr>
<td align="left">sign</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">请求参数的签名</td>
</tr>
<tr>
<td align="left">symbol</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">btc_usd   ltc_usd    eth_usd    etc_usd    bch_usd</td>
</tr>
<tr>
<td align="left">type</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">划转类型。1：币币转合约 2：合约转币币</td>
</tr>
<tr>
<td align="left">amount</td>
<td align="left">String</td>
<td align="left">是</td>
<td align="left">划转币的数量</td>
</tr></tbody></table>
