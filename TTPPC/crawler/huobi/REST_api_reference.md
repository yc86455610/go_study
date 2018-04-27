
          <h1>
<a id="user-content-api-reference" class="anchor" href="#api-reference" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>API Reference</h1>
<p><code>请务必在header中设置user agent为 &#39;User-Agent&#39;: &#39;Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36&#39;</code></p>
<p><code>symbol 规则： 基础币种+计价币种。如BTC/USDT，symbol为btcusdt；ETH/BTC， symbol为ethbtc。以此类推</code></p>
<h3>
<a id="user-content-行情api" class="anchor" href="#%E8%A1%8C%E6%83%85api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>行情API</h3>
<h4>
<a id="user-content-get-markethistorykline-获取k线数据" class="anchor" href="#get-markethistorykline-%E8%8E%B7%E5%8F%96k%E7%BA%BF%E6%95%B0%E6%8D%AE" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/history/kline 获取K线数据</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>period</td>
<td>true</td>
<td>string</td>
<td>K线类型</td>
<td></td>
<td>1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year</td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>integer</td>
<td>获取数量</td>
<td>150</td>
<td>[1,2000]</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td>请求处理结果</td>
<td>&#34;ok&#34; , &#34;error&#34;</td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>number</td>
<td>响应生成时间点，单位：毫秒</td>
<td></td>
</tr>
<tr>
<td>tick</td>
<td>true</td>
<td>object</td>
<td>KLine 数据</td>
<td></td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.kline.$period</td>
<td></td>
</tr>
</tbody>
</table>
<p>data 说明:</p>
<pre><code>  &#34;data&#34;: [
{
    &#34;id&#34;: K线id,
    &#34;amount&#34;: 成交量,
    &#34;count&#34;: 成交笔数,
    &#34;open&#34;: 开盘价,
    &#34;close&#34;: 收盘价,当K线为最晚的一根时，是最新成交价
    &#34;low&#34;: 最低价,
    &#34;high&#34;: 最高价,
    &#34;vol&#34;: 成交额, 即 sum(每一笔成交价 * 该笔的成交量)
  }
]
</code></pre>
<p>请求响应示例:</p>
<pre><code>/* GET /market/history/kline?period=1day&amp;size=200&amp;symbol=btcusdt */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;ch&#34;: &#34;market.btcusdt.kline.1day&#34;,
  &#34;ts&#34;: 1499223904680,
  “data”: [
{
    &#34;id&#34;: 1499184000,
    &#34;amount&#34;: 37593.0266,
    &#34;count&#34;: 0,
    &#34;open&#34;: 1935.2000,
    &#34;close&#34;: 1879.0000,
    &#34;low&#34;: 1856.0000,
    &#34;high&#34;: 1940.0000,
    &#34;vol&#34;: 71031537.97866500
  },
// more data here
]
}

/* GET /market/history/kline?period=not-exist&amp;size=200&amp;symbol=ethusdt */
{
  &#34;ts&#34;: 1490758171271,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid period&#34;
}

/* GET /market/history/kline?period=1day&amp;size=not-exist&amp;symbol=ethusdt */
{
  &#34;ts&#34;: 1490758221221,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;bad-request&#34;,
  &#34;err-msg&#34;: &#34;invalid size, valid range: [1,2000]&#34;
}

/* GET /market/history/kline?period=1day&amp;size=200&amp;symbol=not-exist */
{
  &#34;ts&#34;: 1490758171271,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid symbol&#34;
}
</code></pre>
<h4>
<a id="user-content-get-marketdetailmerged-获取聚合行情ticker" class="anchor" href="#get-marketdetailmerged-%E8%8E%B7%E5%8F%96%E8%81%9A%E5%90%88%E8%A1%8C%E6%83%85ticker" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/detail/merged 获取聚合行情(Ticker)</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td>请求处理结果</td>
<td>&#34;ok&#34; , &#34;error&#34;</td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>number</td>
<td>响应生成时间点，单位：毫秒</td>
<td></td>
</tr>
<tr>
<td>tick</td>
<td>true</td>
<td>object</td>
<td>K线数据</td>
<td></td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.detail.merged</td>
<td></td>
</tr>
</tbody>
</table>
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
    &#34;bid&#34;: [买1价,买1量],
    &#34;ask&#34;: [卖1价,卖1量]
  }

</code></pre>
<p>请求响应示例:</p>
<pre><code>/* GET /market/detail/merged?symbol=ethusdt */
{
&#34;status&#34;:&#34;ok&#34;,
&#34;ch&#34;:&#34;market.ethusdt.detail.merged&#34;,
&#34;ts&#34;:1499225276950,
&#34;tick&#34;:{
  &#34;id&#34;:1499225271,
  &#34;ts&#34;:1499225271000,
  &#34;close&#34;:1885.0000,
  &#34;open&#34;:1960.0000,
  &#34;high&#34;:1985.0000,
  &#34;low&#34;:1856.0000,
  &#34;amount&#34;:81486.2926,
  &#34;count&#34;:42122,
  &#34;vol&#34;:157052744.85708200,
  &#34;ask&#34;:[1885.0000,21.8804],
  &#34;bid&#34;:[1884.0000,1.6702]
  }
}

/* GET /market/detail/merged?symbol=not-exist */
{
  &#34;ts&#34;: 1490758171271,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid symbol”
}

</code></pre>
<h4>
<a id="user-content-get-marketdepth-获取-market-depth-数据" class="anchor" href="#get-marketdepth-%E8%8E%B7%E5%8F%96-market-depth-%E6%95%B0%E6%8D%AE" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/depth 获取 Market Depth 数据</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>Depth 类型</td>
<td></td>
<td>step0, step1, step2, step3, step4, step5（合并深度0-5）；step0时，不合并深度</td>
</tr>
</tbody>
</table>
<ul>
<li>用户选择“合并深度”时，一定报价精度内的市场挂单将予以合并显示。合并深度仅改变显示方式，不改变实际成交价格。</li>
</ul>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td></td>
<td>&#34;ok&#34; 或者 &#34;error&#34;</td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>number</td>
<td>响应生成时间点，单位：毫秒</td>
<td></td>
</tr>
<tr>
<td>tick</td>
<td>true</td>
<td>object</td>
<td>Depth 数据</td>
<td></td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.depth.$type</td>
<td></td>
</tr>
</tbody>
</table>
<p>tick 说明:</p>
<pre><code>  &#34;tick&#34;: {
    &#34;id&#34;: 消息id,
    &#34;ts&#34;: 消息生成时间，单位：毫秒,
    &#34;bids&#34;: 买盘,[price(成交价), amount(成交量)], 按price降序,
    &#34;asks&#34;: 卖盘,[price(成交价), amount(成交量)], 按price升序
  }
</code></pre>
<p>请求响应示例:</p>
<pre><code>/* GET /market/depth?symbol=ethusdt&amp;type=step1 */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;ch&#34;: &#34;market.btcusdt.depth.step1&#34;,
  &#34;ts&#34;: 1489472598812,
  &#34;tick&#34;: {
    &#34;id&#34;: 1489464585407,
    &#34;ts&#34;: 1489464585407,
    &#34;bids&#34;: [
      [7964, 0.0678], // [price, amount]
      [7963, 0.9162],
      [7961, 0.1],
      [7960, 12.8898],
      [7958, 1.2],
      [7955, 2.1009],
      [7954, 0.4708],
      [7953, 0.0564],
      [7951, 2.8031],
      [7950, 13.7785],
      [7949, 0.125],
      [7948, 4],
      [7942, 0.4337],
      [7940, 6.1612],
      [7936, 0.02],
      [7935, 1.3575],
      [7933, 2.002],
      [7932, 1.3449],
      [7930, 10.2974],
      [7929, 3.2226]
    ],
    &#34;asks&#34;: [
      [7979, 0.0736],
      [7980, 1.0292],
      [7981, 5.5652],
      [7986, 0.2416],
      [7990, 1.9970],
      [7995, 0.88],
      [7996, 0.0212],
      [8000, 9.2609],
      [8002, 0.02],
      [8008, 1],
      [8010, 0.8735],
      [8011, 2.36],
      [8012, 0.02],
      [8014, 0.1067],
      [8015, 12.9118],
      [8016, 2.5206],
      [8017, 0.0166],
      [8018, 1.3218],
      [8019, 0.01],
      [8020, 13.6584]
    ]
  }
}

/* GET /market/depth?symbol=ethusdt&amp;type=not-exist */
{
  &#34;ts&#34;: 1490759358099,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid type&#34;
}
</code></pre>
<h4>
<a id="user-content-get-markettrade-获取-trade-detail-数据" class="anchor" href="#get-markettrade-%E8%8E%B7%E5%8F%96-trade-detail-%E6%95%B0%E6%8D%AE" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/trade 获取 Trade Detail 数据</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td></td>
<td>&#34;ok&#34; 或者 &#34;error&#34;</td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>number</td>
<td>响应生成时间点，单位：毫秒</td>
<td></td>
</tr>
<tr>
<td>tick</td>
<td>true</td>
<td>object</td>
<td>Trade 数据</td>
<td></td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.trade.detail</td>
<td></td>
</tr>
</tbody>
</table>
<p>tick 说明：</p>
<pre><code>  &#34;tick&#34;: {
    &#34;id&#34;: 消息id,
    &#34;ts&#34;: 最新成交时间,
    &#34;data&#34;: [
      {
        &#34;id&#34;: 成交id,
        &#34;price&#34;: 成交价钱,
        &#34;amount&#34;: 成交量,
        &#34;direction&#34;: 主动成交方向,
        &#34;ts&#34;: 成交时间
      }
    ]
  }
</code></pre>
<p>请求响应例子:</p>
<pre><code>/* GET /market/trade?symbol=ethusdt */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;ch&#34;: &#34;market.btcusdt.trade.detail&#34;,
  &#34;ts&#34;: 1489473346905,
  &#34;tick&#34;: {
    &#34;id&#34;: 600848670,
    &#34;ts&#34;: 1489464451000,
    &#34;data&#34;: [
      {
        &#34;id&#34;: 600848670,
        &#34;price&#34;: 7962.62,
        &#34;amount&#34;: 0.0122,
        &#34;direction&#34;: &#34;buy&#34;,
        &#34;ts&#34;: 1489464451000
      }
    ]
  }
}

/* GET /market/trade?symbol=not-exist */
{
  &#34;ts&#34;: 1490759506429,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid symbol&#34;
}
</code></pre>
<h4>
<a id="user-content-get-markethistorytrade-批量获取最近的交易记录" class="anchor" href="#get-markethistorytrade-%E6%89%B9%E9%87%8F%E8%8E%B7%E5%8F%96%E6%9C%80%E8%BF%91%E7%9A%84%E4%BA%A4%E6%98%93%E8%AE%B0%E5%BD%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/history/trade 批量获取最近的交易记录</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>integer</td>
<td>获取交易记录的数量</td>
<td>1</td>
<td>[1, 2000]</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td></td>
<td></td>
<td>ok, error</td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.trade.detail</td>
<td></td>
<td></td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>integer</td>
<td>发送时间</td>
<td></td>
<td></td>
</tr>
<tr>
<td>data</td>
<td>true</td>
<td>object</td>
<td>成交记录</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>data 说明：</p>
<pre><code>  &#34;data&#34;: {
    &#34;id&#34;: 消息id,
    &#34;ts&#34;: 最新成交时间,
    &#34;data&#34;: [
      {
        &#34;id&#34;: 成交id,
        &#34;price&#34;: 成交价,
        &#34;amount&#34;: 成交量,
        &#34;direction&#34;: 主动成交方向,
        &#34;ts&#34;: 成交时间
      }
    ]
  }
</code></pre>
<p>请求响应例子:</p>
<pre><code>/* GET /market/history/trade?symbol=ethusdt */
{
    &#34;status&#34;: &#34;ok&#34;,
    &#34;ch&#34;: &#34;market.ethusdt.trade.detail&#34;,
    &#34;ts&#34;: 1502448925216,
    &#34;data&#34;: [
        {
            &#34;id&#34;: 31459998,
            &#34;ts&#34;: 1502448920106,
            &#34;data&#34;: [
                {
                    &#34;id&#34;: 17592256642623,
                    &#34;amount&#34;: 0.04,
                    &#34;price&#34;: 1997,
                    &#34;direction&#34;: &#34;buy&#34;,
                    &#34;ts&#34;: 1502448920106
                }
            ]
        }
    ]
}
</code></pre>
<h4>
<a id="user-content-get-marketdetail-获取-market-detail-24小时成交量数据" class="anchor" href="#get-marketdetail-%E8%8E%B7%E5%8F%96-market-detail-24%E5%B0%8F%E6%97%B6%E6%88%90%E4%BA%A4%E9%87%8F%E6%95%B0%E6%8D%AE" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /market/detail 获取 Market Detail 24小时成交量数据</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>status</td>
<td>true</td>
<td>string</td>
<td></td>
<td>&#34;ok&#34; 或者 &#34;error&#34;</td>
</tr>
<tr>
<td>ts</td>
<td>true</td>
<td>number</td>
<td>响应生成时间点，单位：毫秒</td>
<td></td>
</tr>
<tr>
<td>tick</td>
<td>true</td>
<td>object</td>
<td>Detail 数据</td>
<td></td>
</tr>
<tr>
<td>ch</td>
<td>true</td>
<td>string</td>
<td>数据所属的 channel，格式： market.$symbol.depth.$type</td>
<td></td>
</tr>
</tbody>
</table>
<p>tick 说明:</p>
<pre><code>  &#34;tick&#34;: {
    &#34;id&#34;: 消息id,
    &#34;ts&#34;: 24小时统计时间,
    &#34;amount&#34;: 24小时成交量,
    &#34;open&#34;: 前推24小时成交价,
    &#34;close&#34;: 当前成交价,
    &#34;high&#34;: 近24小时最高价,
    &#34;low&#34;: 近24小时最低价,
    &#34;count&#34;: 近24小时累积成交数,
    &#34;vol&#34;: 近24小时累积成交额, 即 sum(每一笔成交价 * 该笔的成交量)
  }
</code></pre>
<p>请求响应例子</p>
<pre><code>/* GET /market/detail?symbol=ethusdt */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;ch&#34;: &#34;market.btcusdt.detail&#34;,
  &#34;ts&#34;: 1489473538996,
  &#34;tick&#34;: {
    &#34;amount&#34;: 4316.4346,
    &#34;open&#34;: 8090.54,
    &#34;close&#34;: 7962.62,
    &#34;high&#34;: 8119.00,
    &#34;ts&#34;: 1489464451000,
    &#34;id&#34;: 1489464451,
    &#34;count&#34;: 9595,
    &#34;low&#34;: 7875.00,
    &#34;vol&#34;: 34497276.905760
  }
}

/* GET /market/detail?symbol=not-exists */
{
  &#34;ts&#34;: 1490759594752,
  &#34;status&#34;: &#34;error&#34;,
  &#34;err-code&#34;: &#34;invalid-parameter&#34;,
  &#34;err-msg&#34;: &#34;invalid symbol&#34;
}
</code></pre>
<h3>
<a id="user-content-公共api" class="anchor" href="#%E5%85%AC%E5%85%B1api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>公共API</h3>
<h4>
<a id="user-content-get-v1commonsymbols-查询pro站支持的所有交易对及精度" class="anchor" href="#get-v1commonsymbols-%E6%9F%A5%E8%AF%A2pro%E7%AB%99%E6%94%AF%E6%8C%81%E7%9A%84%E6%89%80%E6%9C%89%E4%BA%A4%E6%98%93%E5%AF%B9%E5%8F%8A%E7%B2%BE%E5%BA%A6" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/common/symbols 查询Pro站支持的所有交易对及精度</h4>
<h4>
<a id="user-content-get-v1hadaxcommonsymbols-查询hadax站支持的所有交易对及精度" class="anchor" href="#get-v1hadaxcommonsymbols-%E6%9F%A5%E8%AF%A2hadax%E7%AB%99%E6%94%AF%E6%8C%81%E7%9A%84%E6%89%80%E6%9C%89%E4%BA%A4%E6%98%93%E5%AF%B9%E5%8F%8A%E7%B2%BE%E5%BA%A6" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/hadax/common/symbols 查询HADAX站支持的所有交易对及精度</h4>
<p>请求参数:
(无)</p>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>base-currency</td>
<td>true</td>
<td>string</td>
<td>基础币种</td>
<td></td>
</tr>
<tr>
<td>quote-currency</td>
<td>true</td>
<td>string</td>
<td>计价币种</td>
<td></td>
</tr>
<tr>
<td>price-precision</td>
<td>true</td>
<td>string</td>
<td>价格精度位数（0为个位）</td>
<td></td>
</tr>
<tr>
<td>amount-precision</td>
<td>true</td>
<td>string</td>
<td>数量精度位数（0为个位）</td>
<td></td>
</tr>
<tr>
<td>symbol-partition</td>
<td>true</td>
<td>string</td>
<td>交易区</td>
<td>main主区，innovation创新区，bifurcation分叉区</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/common/symbols */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;base-currency&#34;: &#34;eth&#34;,
      &#34;quote-currency&#34;: &#34;usdt&#34;,
      &#34;symbol&#34;: &#34;ethusdt&#34;
    }
    {
      &#34;base-currency&#34;: &#34;etc&#34;,
      &#34;quote-currency&#34;: &#34;usdt&#34;,
      &#34;symbol&#34;: &#34;etcusdt&#34;
    }
  ]
}
</code></pre>
<h4>
<a id="user-content-get-v1commoncurrencys-查询pro站支持的所有币种" class="anchor" href="#get-v1commoncurrencys-%E6%9F%A5%E8%AF%A2pro%E7%AB%99%E6%94%AF%E6%8C%81%E7%9A%84%E6%89%80%E6%9C%89%E5%B8%81%E7%A7%8D" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/common/currencys 查询Pro站支持的所有币种</h4>
<h4>
<a id="user-content-get-v1hadaxcommoncurrencys-查询hadax站支持的所有币种" class="anchor" href="#get-v1hadaxcommoncurrencys-%E6%9F%A5%E8%AF%A2hadax%E7%AB%99%E6%94%AF%E6%8C%81%E7%9A%84%E6%89%80%E6%9C%89%E5%B8%81%E7%A7%8D" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/hadax/common/currencys 查询HADAX站支持的所有币种</h4>
<p>请求参数:</p>
<p>(无)</p>
<p>响应数据:</p>
<pre><code>currency list
</code></pre>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/common/currencys */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    &#34;usdt&#34;,
    &#34;eth&#34;,
    &#34;etc&#34;
  ]
}
</code></pre>
<h4>
<a id="user-content-get-v1commontimestamp-查询系统当前时间" class="anchor" href="#get-v1commontimestamp-%E6%9F%A5%E8%AF%A2%E7%B3%BB%E7%BB%9F%E5%BD%93%E5%89%8D%E6%97%B6%E9%97%B4" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/common/timestamp 查询系统当前时间</h4>
<p>请求参数:</p>
<p>(无)</p>
<p>响应数据:</p>
<pre><code>系统时间戳
</code></pre>
<p>请求响应例子</p>
<pre><code>/* GET /v1/common/timestamp */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 1494900087029
}
</code></pre>
<h3>
<a id="user-content-用户资产api" class="anchor" href="#%E7%94%A8%E6%88%B7%E8%B5%84%E4%BA%A7api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>用户资产API</h3>
<h4>
<a id="user-content-get-v1accountaccounts-查询当前用户的所有账户即account-idpro站和hadax-account-id通用" class="anchor" href="#get-v1accountaccounts-%E6%9F%A5%E8%AF%A2%E5%BD%93%E5%89%8D%E7%94%A8%E6%88%B7%E7%9A%84%E6%89%80%E6%9C%89%E8%B4%A6%E6%88%B7%E5%8D%B3account-idpro%E7%AB%99%E5%92%8Chadax-account-id%E9%80%9A%E7%94%A8" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/account/accounts 查询当前用户的所有账户(即account-id)，Pro站和HADAX account-id通用</h4>
<p>请求参数:</p>
<p>无</p>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>account-id</td>
<td></td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>账户状态</td>
<td>working：正常, lock：账户被锁定</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>账户类型</td>
<td>spot：现货账户</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/account/accounts */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;id&#34;: 100009,
      &#34;type&#34;: &#34;spot&#34;,
      &#34;state&#34;: &#34;working&#34;,
      &#34;user-id&#34;: 1000
    }
  ]
}
</code></pre>
<h4>
<a id="user-content-get-v1accountaccountsaccount-idbalance-查询pro站指定账户的余额" class="anchor" href="#get-v1accountaccountsaccount-idbalance-%E6%9F%A5%E8%AF%A2pro%E7%AB%99%E6%8C%87%E5%AE%9A%E8%B4%A6%E6%88%B7%E7%9A%84%E4%BD%99%E9%A2%9D" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/account/accounts/{account-id}/balance 查询Pro站指定账户的余额</h4>
<h4>
<a id="user-content-get-v1hadaxaccountaccountsaccount-idbalance--查询hadax站指定账户的余额" class="anchor" href="#get-v1hadaxaccountaccountsaccount-idbalance--%E6%9F%A5%E8%AF%A2hadax%E7%AB%99%E6%8C%87%E5%AE%9A%E8%B4%A6%E6%88%B7%E7%9A%84%E4%BD%99%E9%A2%9D" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/hadax/account/accounts/{account-id}/balance  查询HADAX站指定账户的余额</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>account-id</td>
<td>true</td>
<td>string</td>
<td>account-id，填在 path 中，可用 GET /v1/account/accounts 获取</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<ul>
<li>如果不知道自己的账户ID，请使用 <code>GET /v1/account/accounts</code> 查询</li>
</ul>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>账户 ID</td>
<td></td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>账户状态</td>
<td>working：正常  lock：账户被锁定</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>账户类型</td>
<td>spot：现货账户</td>
</tr>
<tr>
<td>list</td>
<td>false</td>
<td>Array</td>
<td>子账户数组</td>
<td></td>
</tr>
</tbody>
</table>
<p>list字段说明</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>balance</td>
<td>true</td>
<td>string</td>
<td>余额</td>
<td></td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>类型</td>
<td>trade: 交易余额，frozen: 冻结余额</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/account/accounts/{account-id}/balance */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: {
    &#34;id&#34;: 100009,
    &#34;type&#34;: &#34;spot&#34;,
    &#34;state&#34;: &#34;working&#34;,
    &#34;list&#34;: [
      {
        &#34;currency&#34;: &#34;usdt&#34;,
        &#34;type&#34;: &#34;trade&#34;,
        &#34;balance&#34;: &#34;500009195917.4362872650&#34;
      },
      {
        &#34;currency&#34;: &#34;usdt&#34;,
        &#34;type&#34;: &#34;frozen&#34;,
        &#34;balance&#34;: &#34;328048.1199920000&#34;
      },
     {
        &#34;currency&#34;: &#34;etc&#34;,
        &#34;type&#34;: &#34;trade&#34;,
        &#34;balance&#34;: &#34;499999894616.1302471000&#34;
      },
      {
        &#34;currency&#34;: &#34;etc&#34;,
        &#34;type&#34;: &#34;frozen&#34;,
        &#34;balance&#34;: &#34;9786.6783000000&#34;
      }
     {
        &#34;currency&#34;: &#34;eth&#34;,
        &#34;type&#34;: &#34;trade&#34;,
        &#34;balance&#34;: &#34;499999894616.1302471000&#34;
      },
      {
        &#34;currency&#34;: &#34;eth”,
        &#34;type&#34;: &#34;frozen&#34;,
        &#34;balance&#34;: &#34;9786.6783000000&#34;
      }
    ],
    &#34;user-id&#34;: 1000
  }
}
</code></pre>
<h2>
<a id="user-content-交易api" class="anchor" href="#%E4%BA%A4%E6%98%93api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>交易API</h2>
<h4>
<a id="user-content-post-v1orderordersplace-pro站下单" class="anchor" href="#post-v1orderordersplace-pro%E7%AB%99%E4%B8%8B%E5%8D%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/order/orders/place Pro站下单</h4>
<h4>
<a id="user-content-post-v1hadaxorderordersplace-hadax站下单" class="anchor" href="#post-v1hadaxorderordersplace-hadax%E7%AB%99%E4%B8%8B%E5%8D%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/hadax/order/orders/place HADAX站下单</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>account-id</td>
<td>true</td>
<td>string</td>
<td>账户 ID，使用accounts方法获得。币币交易使用‘spot’账户的accountid；借贷资产交易，请使用‘margin’账户的accountid</td>
<td></td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>限价单表示下单数量，市价买单时表示买多少钱，市价卖单时表示卖多少币</td>
<td></td>
<td></td>
</tr>
<tr>
<td>price</td>
<td>false</td>
<td>string</td>
<td>下单价格，市价单不传该参数</td>
<td></td>
<td></td>
</tr>
<tr>
<td>source</td>
<td>false</td>
<td>string</td>
<td>订单来源</td>
<td>api，如果使用借贷资产交易，请填写‘margin-api’</td>
<td></td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>订单类型</td>
<td></td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>false</td>
<td>string</td>
<td>订单ID</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/order/orders/place {
   &#34;account-id&#34;: &#34;100009&#34;,
   &#34;amount&#34;: &#34;10.1&#34;,
   &#34;price&#34;: &#34;100.1&#34;,
   &#34;source&#34;: &#34;api&#34;,
   &#34;symbol&#34;: &#34;ethusdt&#34;,
   &#34;type&#34;: &#34;buy-limit&#34;
 } */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: &#34;59378&#34;
}
</code></pre>
<h4>
<a id="user-content-post-v1orderordersorder-idsubmitcancel--申请撤销一个订单请求" class="anchor" href="#post-v1orderordersorder-idsubmitcancel--%E7%94%B3%E8%AF%B7%E6%92%A4%E9%94%80%E4%B8%80%E4%B8%AA%E8%AE%A2%E5%8D%95%E8%AF%B7%E6%B1%82" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/order/orders/{order-id}/submitcancel  申请撤销一个订单请求</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>order-id</td>
<td>true</td>
<td>string</td>
<td>订单ID，填在path中</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>true</td>
<td>string</td>
<td>订单 ID</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/order/orders/{order-id}/submitcancel */
{
  &#34;status&#34;: &#34;ok&#34;,//注意，返回OK表示撤单请求成功。订单是否撤销成功请调用订单查询接口查询该订单状态
  &#34;data&#34;: &#34;59378&#34;
}
</code></pre>
<h4>
<a id="user-content-post-v1orderordersbatchcancel-批量撤销订单" class="anchor" href="#post-v1orderordersbatchcancel-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E8%AE%A2%E5%8D%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/order/orders/batchcancel 批量撤销订单</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>order-ids</td>
<td>true</td>
<td>list</td>
<td>撤销订单ID列表</td>
<td></td>
<td>单次不超过50个订单id</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>false</td>
<td>map</td>
<td>撤单结果</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/order/orders/batchcancel */
{
  &#34;order-ids&#34;: [
    &#34;1&#34;, &#34;2&#34;, &#34;3&#34;
  ]
}

-----

{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: {
    &#34;success&#34;: [
      &#34;1&#34;,
      &#34;3&#34;
    ],
    &#34;failed&#34;: [
      {
        &#34;err-msg&#34;: &#34;记录无效&#34;,
        &#34;order-id&#34;: &#34;2&#34;,
        &#34;err-code&#34;: &#34;base-record-invalid&#34;
      }
    ]
  }
}
</code></pre>
<h4>
<a id="user-content-get-v1orderordersorder-id-查询某个订单详情" class="anchor" href="#get-v1orderordersorder-id-%E6%9F%A5%E8%AF%A2%E6%9F%90%E4%B8%AA%E8%AE%A2%E5%8D%95%E8%AF%A6%E6%83%85" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/order/orders/{order-id} 查询某个订单详情</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>order-id</td>
<td>true</td>
<td>string</td>
<td>订单ID，填在path中</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>account-id</td>
<td>true</td>
<td>long</td>
<td>账户 ID</td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>订单数量</td>
<td></td>
</tr>
<tr>
<td>canceled-at</td>
<td>false</td>
<td>long</td>
<td>订单撤销时间</td>
<td></td>
</tr>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>订单创建时间</td>
<td></td>
</tr>
<tr>
<td>field-amount</td>
<td>true</td>
<td>string</td>
<td>已成交数量</td>
<td></td>
</tr>
<tr>
<td>field-cash-amount</td>
<td>true</td>
<td>string</td>
<td>已成交总金额</td>
<td></td>
</tr>
<tr>
<td>field-fees</td>
<td>true</td>
<td>string</td>
<td>已成交手续费（买入为币，卖出为钱）</td>
<td></td>
</tr>
<tr>
<td>finished-at</td>
<td>false</td>
<td>long</td>
<td>最后成交时间</td>
<td></td>
</tr>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>订单ID</td>
<td></td>
</tr>
<tr>
<td>price</td>
<td>true</td>
<td>string</td>
<td>订单价格</td>
<td></td>
</tr>
<tr>
<td>source</td>
<td>true</td>
<td>string</td>
<td>订单来源</td>
<td>api</td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>订单状态</td>
<td>pre-submitted 准备提交, submitting , submitted 已提交, partial-filled 部分成交, partial-canceled 部分成交撤销, filled 完全成交, canceled 已撤销</td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>订单类型</td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/order/orders/{order-id} */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: {
    &#34;id&#34;: 59378,
    &#34;symbol&#34;: &#34;ethusdt&#34;,
    &#34;account-id&#34;: 100009,
    &#34;amount&#34;: &#34;10.1000000000&#34;,
    &#34;price&#34;: &#34;100.1000000000&#34;,
    &#34;created-at&#34;: 1494901162595,
    &#34;type&#34;: &#34;buy-limit&#34;,
    &#34;field-amount&#34;: &#34;10.1000000000&#34;,
    &#34;field-cash-amount&#34;: &#34;1011.0100000000&#34;,
    &#34;field-fees&#34;: &#34;0.0202000000&#34;,
    &#34;finished-at&#34;: 1494901400468,
    &#34;user-id&#34;: 1000,
    &#34;source&#34;: &#34;api&#34;,
    &#34;state&#34;: &#34;filled&#34;,
    &#34;canceled-at&#34;: 0,
    &#34;exchange&#34;: &#34;huobi&#34;,
    &#34;batch&#34;: &#34;&#34;
  }
}
</code></pre>
<h4>
<a id="user-content-get-v1orderordersorder-idmatchresults--查询某个订单的成交明细" class="anchor" href="#get-v1orderordersorder-idmatchresults--%E6%9F%A5%E8%AF%A2%E6%9F%90%E4%B8%AA%E8%AE%A2%E5%8D%95%E7%9A%84%E6%88%90%E4%BA%A4%E6%98%8E%E7%BB%86" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/order/orders/{order-id}/matchresults  查询某个订单的成交明细</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>order-id</td>
<td>true</td>
<td>string</td>
<td>订单ID，填在path中</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>成交时间</td>
<td></td>
</tr>
<tr>
<td>filled-amount</td>
<td>true</td>
<td>string</td>
<td>成交数量</td>
<td></td>
</tr>
<tr>
<td>filled-fees</td>
<td>true</td>
<td>string</td>
<td>成交手续费</td>
<td></td>
</tr>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>订单成交记录ID</td>
<td></td>
</tr>
<tr>
<td>match-id</td>
<td>true</td>
<td>long</td>
<td>撮合ID</td>
<td></td>
</tr>
<tr>
<td>order-id</td>
<td>true</td>
<td>long</td>
<td>订单 ID</td>
<td></td>
</tr>
<tr>
<td>price</td>
<td>true</td>
<td>string</td>
<td>成交价格</td>
<td></td>
</tr>
<tr>
<td>source</td>
<td>true</td>
<td>string</td>
<td>订单来源</td>
<td>api</td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>订单类型</td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/order/orders/{order-id}/matchresults */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;id&#34;: 29553,
      &#34;order-id&#34;: 59378,
      &#34;match-id&#34;: 59335,
      &#34;symbol&#34;: &#34;ethusdt&#34;,
      &#34;type&#34;: &#34;buy-limit&#34;,
      &#34;source&#34;: &#34;api&#34;,
      &#34;price&#34;: &#34;100.1000000000&#34;,
      &#34;filled-amount&#34;: &#34;9.1155000000&#34;,
      &#34;filled-fees&#34;: &#34;0.0182310000&#34;,
      &#34;created-at&#34;: 1494901400435
    }
  ]
}
</code></pre>
<h4>
<a id="user-content-get-v1orderorders-查询当前委托历史委托" class="anchor" href="#get-v1orderorders-%E6%9F%A5%E8%AF%A2%E5%BD%93%E5%89%8D%E5%A7%94%E6%89%98%E5%8E%86%E5%8F%B2%E5%A7%94%E6%89%98" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/order/orders 查询当前委托、历史委托</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>types</td>
<td>false</td>
<td>string</td>
<td>查询的订单类型组合，使用&#39;,&#39;分割</td>
<td></td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
<tr>
<td>start-date</td>
<td>false</td>
<td>string</td>
<td>查询开始日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>end-date</td>
<td>false</td>
<td>string</td>
<td>查询结束日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>states</td>
<td>true</td>
<td>string</td>
<td>查询的订单状态组合，使用&#39;,&#39;分割</td>
<td></td>
<td>pre-submitted 准备提交, submitted 已提交, partial-filled 部分成交, partial-canceled 部分成交撤销, filled 完全成交, canceled 已撤销</td>
</tr>
<tr>
<td>from</td>
<td>false</td>
<td>string</td>
<td>查询起始 ID</td>
<td></td>
<td></td>
</tr>
<tr>
<td>direct</td>
<td>false</td>
<td>string</td>
<td>查询方向</td>
<td></td>
<td>prev 向前，next 向后</td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>string</td>
<td>查询记录大小</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>account-id</td>
<td>true</td>
<td>long</td>
<td>账户 ID</td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>订单数量</td>
<td></td>
</tr>
<tr>
<td>canceled-at</td>
<td>false</td>
<td>long</td>
<td>订单撤销时间</td>
<td></td>
</tr>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>订单创建时间</td>
<td></td>
</tr>
<tr>
<td>field-amount</td>
<td>true</td>
<td>string</td>
<td>已成交数量</td>
<td></td>
</tr>
<tr>
<td>field-cash-amount</td>
<td>true</td>
<td>string</td>
<td>已成交总金额</td>
<td></td>
</tr>
<tr>
<td>field-fees</td>
<td>true</td>
<td>string</td>
<td>已成交手续费（买入为币，卖出为钱）</td>
<td></td>
</tr>
<tr>
<td>finished-at</td>
<td>false</td>
<td>long</td>
<td>最后成交时间</td>
<td></td>
</tr>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>订单ID</td>
<td></td>
</tr>
<tr>
<td>price</td>
<td>true</td>
<td>string</td>
<td>订单价格</td>
<td></td>
</tr>
<tr>
<td>source</td>
<td>true</td>
<td>string</td>
<td>订单来源</td>
<td>api</td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>订单状态</td>
<td>pre-submitted 准备提交, submitting , submitted 已提交, partial-filled 部分成交, partial-canceled 部分成交撤销, filled 完全成交, canceled 已撤销</td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>订单类型</td>
<td>submit-cancel：已提交撤单申请  ,buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/order/orders */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;id&#34;: 59378,
      &#34;symbol&#34;: &#34;ethusdt&#34;,
      &#34;account-id&#34;: 100009,
      &#34;amount&#34;: &#34;10.1000000000&#34;,
      &#34;price&#34;: &#34;100.1000000000&#34;,
      &#34;created-at&#34;: 1494901162595,
      &#34;type&#34;: &#34;buy-limit&#34;,
      &#34;field-amount&#34;: &#34;10.1000000000&#34;,
      &#34;field-cash-amount&#34;: &#34;1011.0100000000&#34;,
      &#34;field-fees&#34;: &#34;0.0202000000&#34;,
      &#34;finished-at&#34;: 1494901400468,
      &#34;user-id&#34;: 1000,
      &#34;source&#34;: &#34;api&#34;,
      &#34;state&#34;: &#34;filled&#34;,
      &#34;canceled-at&#34;: 0,
      &#34;exchange&#34;: &#34;huobi&#34;,
      &#34;batch&#34;: &#34;&#34;
    }
  ]
}
</code></pre>
<h4>
<a id="user-content-get-v1ordermatchresults-查询当前成交历史成交" class="anchor" href="#get-v1ordermatchresults-%E6%9F%A5%E8%AF%A2%E5%BD%93%E5%89%8D%E6%88%90%E4%BA%A4%E5%8E%86%E5%8F%B2%E6%88%90%E4%BA%A4" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/order/matchresults 查询当前成交、历史成交</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td>btcusdt, bchbtc, rcneth ...</td>
<td></td>
</tr>
<tr>
<td>types</td>
<td>false</td>
<td>string</td>
<td>查询的订单类型组合，使用&#39;,&#39;分割</td>
<td></td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
<tr>
<td>start-date</td>
<td>false</td>
<td>string</td>
<td>查询开始日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>end-date</td>
<td>false</td>
<td>string</td>
<td>查询结束日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>from</td>
<td>false</td>
<td>string</td>
<td>查询起始 ID</td>
<td></td>
<td></td>
</tr>
<tr>
<td>direct</td>
<td>false</td>
<td>string</td>
<td>查询方向</td>
<td></td>
<td>prev 向前，next 向后</td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>string</td>
<td>查询记录大小</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>成交时间</td>
<td></td>
</tr>
<tr>
<td>filled-amount</td>
<td>true</td>
<td>string</td>
<td>成交数量</td>
<td></td>
</tr>
<tr>
<td>filled-fees</td>
<td>true</td>
<td>string</td>
<td>成交手续费</td>
<td></td>
</tr>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>订单成交记录ID</td>
<td></td>
</tr>
<tr>
<td>match-id</td>
<td>true</td>
<td>long</td>
<td>撮合ID</td>
<td></td>
</tr>
<tr>
<td>order-id</td>
<td>true</td>
<td>long</td>
<td>订单 ID</td>
<td></td>
</tr>
<tr>
<td>price</td>
<td>true</td>
<td>string</td>
<td>成交价格</td>
<td></td>
</tr>
<tr>
<td>source</td>
<td>true</td>
<td>string</td>
<td>订单来源</td>
<td>api</td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td>btcusdt, bchbtc, rcneth ...</td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>订单类型</td>
<td>buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/orders/matchresults */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;id&#34;: 29555,
      &#34;order-id&#34;: 59378,
      &#34;match-id&#34;: 59335,
      &#34;symbol&#34;: &#34;ethusdt&#34;,
      &#34;type&#34;: &#34;buy-limit&#34;,
      &#34;source&#34;: &#34;api&#34;,
      &#34;price&#34;: &#34;100.1000000000&#34;,
      &#34;filled-amount&#34;: &#34;0.9845000000&#34;,
      &#34;filled-fees&#34;: &#34;0.0019690000&#34;,
      &#34;created-at&#34;: 1494901400487
    }
  ]
}
</code></pre>
<h2>
<a id="user-content-借贷交易api-重要如果使用借贷资产交易请在下单接口v1orderordersplace请求参数source中填写margin-api" class="anchor" href="#%E5%80%9F%E8%B4%B7%E4%BA%A4%E6%98%93api-%E9%87%8D%E8%A6%81%E5%A6%82%E6%9E%9C%E4%BD%BF%E7%94%A8%E5%80%9F%E8%B4%B7%E8%B5%84%E4%BA%A7%E4%BA%A4%E6%98%93%E8%AF%B7%E5%9C%A8%E4%B8%8B%E5%8D%95%E6%8E%A5%E5%8F%A3v1orderordersplace%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0source%E4%B8%AD%E5%A1%AB%E5%86%99margin-api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>借贷交易API （重要：如果使用借贷资产交易，请在下单接口/v1/order/orders/place请求参数source中填写‘margin-api’）</h2>
<p><code>目前仅支持 USDT 交易区和 BTC 交易区部分交易对</code></p>
<h4>
<a id="user-content-post-v1dwtransfer-inmargin--现货账户划入至借贷账户" class="anchor" href="#post-v1dwtransfer-inmargin--%E7%8E%B0%E8%B4%A7%E8%B4%A6%E6%88%B7%E5%88%92%E5%85%A5%E8%87%B3%E5%80%9F%E8%B4%B7%E8%B4%A6%E6%88%B7" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/dw/transfer-in/margin  现货账户划入至借贷账户</h4>
<h4>
<a id="user-content-post-v1dwtransfer-outmargin--借贷账户划出至现货账户" class="anchor" href="#post-v1dwtransfer-outmargin--%E5%80%9F%E8%B4%B7%E8%B4%A6%E6%88%B7%E5%88%92%E5%87%BA%E8%87%B3%E7%8E%B0%E8%B4%A7%E8%B4%A6%E6%88%B7" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/dw/transfer-out/margin  借贷账户划出至现货账户</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td></td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>金额</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>true</td>
<td>long</td>
<td>划转ID</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/dw/transfer-in/margin 
{
  &#34;symbol&#34;: &#34;ethusdt&#34;,
  &#34;currency&#34;: &#34;eth&#34;,
  &#34;amount&#34;: &#34;1.0&#34;
} */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 1000
}
</code></pre>
<h4>
<a id="user-content-post-v1marginorders-申请借贷" class="anchor" href="#post-v1marginorders-%E7%94%B3%E8%AF%B7%E5%80%9F%E8%B4%B7" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/margin/orders 申请借贷</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td></td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>金额</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>true</td>
<td>long</td>
<td>订单号</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/margin/orders {
   &#34;amount&#34;: &#34;10.1&#34;,
   &#34;symbol&#34;: &#34;ethusdt&#34;,
   &#34;currency&#34;: &#34;eth&#34;
} */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 59378
}
</code></pre>
<h4>
<a id="user-content-post-v1marginordersorder-idrepay-归还借贷" class="anchor" href="#post-v1marginordersorder-idrepay-%E5%BD%92%E8%BF%98%E5%80%9F%E8%B4%B7" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/margin/orders/{order-id}/repay 归还借贷</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>order-id</td>
<td>true</td>
<td>long</td>
<td>借贷订单 ID，写在path中</td>
<td></td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>还款量</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>true</td>
<td>long</td>
<td>订单号</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/margin/orders/59378/repay {
   &#34;amount&#34;: &#34;10.1&#34;
}*/
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 59378
}
</code></pre>
<h4>
<a id="user-content-get-v1marginloan-orders--借贷订单" class="anchor" href="#get-v1marginloan-orders--%E5%80%9F%E8%B4%B7%E8%AE%A2%E5%8D%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/margin/loan-orders  借贷订单</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
<td></td>
</tr>
<tr>
<td>start-date</td>
<td>false</td>
<td>string</td>
<td>查询开始日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>end-date</td>
<td>false</td>
<td>string</td>
<td>查询结束日期, 日期格式yyyy-mm-dd</td>
<td></td>
<td></td>
</tr>
<tr>
<td>states</td>
<td>false</td>
<td>string</td>
<td>状态</td>
<td></td>
<td></td>
</tr>
<tr>
<td>from</td>
<td>false</td>
<td>string</td>
<td>查询起始 ID</td>
<td></td>
<td></td>
</tr>
<tr>
<td>direct</td>
<td>false</td>
<td>string</td>
<td>查询方向</td>
<td></td>
<td>prev 向前，next 向后</td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>string</td>
<td>查询记录大小</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td>订单号</td>
<td></td>
</tr>
<tr>
<td>user-id</td>
<td>true</td>
<td>long</td>
<td>用户ID</td>
<td></td>
</tr>
<tr>
<td>account-id</td>
<td>true</td>
<td>long</td>
<td>账户ID</td>
<td></td>
</tr>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
</tr>
<tr>
<td>loan-amount</td>
<td>true</td>
<td>string</td>
<td>借贷本金总额</td>
<td></td>
</tr>
<tr>
<td>loan-balance</td>
<td>true</td>
<td>string</td>
<td>未还本金</td>
<td></td>
</tr>
<tr>
<td>interest-rate</td>
<td>true</td>
<td>string</td>
<td>利率</td>
<td></td>
</tr>
<tr>
<td>interest-amount</td>
<td>true</td>
<td>string</td>
<td>利息总额</td>
<td></td>
</tr>
<tr>
<td>interest-balance</td>
<td>true</td>
<td>string</td>
<td>未还利息</td>
<td></td>
</tr>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>借贷发起时间</td>
<td></td>
</tr>
<tr>
<td>accrued-at</td>
<td>true</td>
<td>long</td>
<td>最近一次计息时间</td>
<td></td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>订单状态</td>
<td>created 未放款，accrual 已放款，cleared 已还清，invalid 异常</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/margin/loan-orders?symbol=btcusdt 

*/
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: [
    {
      &#34;loan-balance&#34;: &#34;0.100000000000000000&#34;,
      &#34;interest-balance&#34;: &#34;0.000200000000000000&#34;,
      &#34;interest-rate&#34;: &#34;0.002000000000000000&#34;,
      &#34;loan-amount&#34;: &#34;0.100000000000000000&#34;,
      &#34;accrued-at&#34;: 1511169724531,
      &#34;interest-amount&#34;: &#34;0.000200000000000000&#34;,
      &#34;symbol&#34;: &#34;ethbtc&#34;,
      &#34;currency&#34;: &#34;btc&#34;,
      &#34;id&#34;: 394,
      &#34;state&#34;: &#34;accrual&#34;,
      &#34;account-id&#34;: 17747,
      &#34;user-id&#34;: 119913,
      &#34;created-at&#34;: 1511169724531
    }
  ]
}

</code></pre>
<h4>
<a id="user-content-get-v1marginaccountsbalance-借贷账户详情" class="anchor" href="#get-v1marginaccountsbalance-%E5%80%9F%E8%B4%B7%E8%B4%A6%E6%88%B7%E8%AF%A6%E6%83%85" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/margin/accounts/balance 借贷账户详情</h4>
<p>请求参数</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>false</td>
<td>string</td>
<td>交易对，作为get参数</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>symbol</td>
<td>true</td>
<td>string</td>
<td>交易对</td>
<td></td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>账户状态</td>
<td>working,fl-sys,fl-mgt,fl-end</td>
</tr>
<tr>
<td>risk-rate</td>
<td>true</td>
<td>object</td>
<td>风险率</td>
<td></td>
</tr>
<tr>
<td>fl-price</td>
<td>true</td>
<td>string</td>
<td>爆仓价</td>
<td></td>
</tr>
<tr>
<td>list</td>
<td>true</td>
<td>array</td>
<td>子账户列表</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/margin/accounts/balance?symbol=btcusdt

*/
{
    &#34;status&#34;: &#34;ok&#34;,
    &#34;data&#34;: [
        {
            &#34;id&#34;: 18264,
            &#34;type&#34;: &#34;margin&#34;,
            &#34;state&#34;: &#34;working&#34;,
            &#34;symbol&#34;: &#34;btcusdt&#34;,
            &#34;fl-price&#34;: &#34;0&#34;,
            &#34;fl-type&#34;: &#34;safe&#34;,
            &#34;risk-rate&#34;: &#34;475.952571086994250554&#34;,
            &#34;list&#34;: [
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;trade&#34;,
                    &#34;balance&#34;: &#34;1168.533000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;frozen&#34;,
                    &#34;balance&#34;: &#34;0.000000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;loan&#34;,
                    &#34;balance&#34;: &#34;-2.433000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;interest&#34;,
                    &#34;balance&#34;: &#34;-0.000533000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;trade&#34;,//借贷账户可用
                    &#34;balance&#34;: &#34;1313.534000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;frozen&#34;,//借贷账户冻结
                    &#34;balance&#34;: &#34;0.000000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;loan&#34;,//已借贷
                    &#34;balance&#34;: &#34;-140.234099999999999920&#34;
                },
                {
                    &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;interest&#34;,//usdt待还利息
                    &#34;balance&#34;: &#34;-0.931206660000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;transfer-out-available&#34;,//可转btc
                    &#34;balance&#34;: &#34;1163.872174670000000000&#34;
                },
                { &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;transfer-out-available&#34;,//可转usdt
                    &#34;balance&#34;: &#34;1313.534000000000000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;btc&#34;,
                    &#34;type&#34;: &#34;loan-available&#34;,//可借btc
                    &#34;balance&#34;: &#34;8161.876538350676000000&#34;
                },
                {
                    &#34;currency&#34;: &#34;usdt&#34;,
                    &#34;type&#34;: &#34;loan-available&#34;,//可借usdt
                    &#34;balance&#34;: &#34;49859.765900000000000080&#34;
                }
            ]
        }
    ]
}

</code></pre>
<h2>
<a id="user-content-虚拟币提现api" class="anchor" href="#%E8%99%9A%E6%8B%9F%E5%B8%81%E6%8F%90%E7%8E%B0api" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>虚拟币提现API</h2>
<blockquote>
<p><strong>仅支持提现到【Pro站提币地址列表中的提币地址】</strong></p>
</blockquote>
<h4>
<a id="user-content-post-v1dwwithdrawapicreate-申请提现虚拟币" class="anchor" href="#post-v1dwwithdrawapicreate-%E7%94%B3%E8%AF%B7%E6%8F%90%E7%8E%B0%E8%99%9A%E6%8B%9F%E5%B8%81" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/dw/withdraw/api/create 申请提现虚拟币</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>address</td>
<td>true</td>
<td>string</td>
<td>提现地址</td>
<td></td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>string</td>
<td>提币数量</td>
<td></td>
<td></td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>资产类型</td>
<td></td>
<td>btc, ltc, bch, eth, etc ...(火币Pro支持的币种)</td>
</tr>
<tr>
<td>fee</td>
<td>false</td>
<td>string</td>
<td>转账手续费</td>
<td></td>
<td></td>
</tr>
<tr>
<td>addr-tag</td>
<td>false</td>
<td>string</td>
<td>虚拟币共享地址tag，XRP特有</td>
<td></td>
<td>格式, &#34;123&#34;类的整数字符串</td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>false</td>
<td>long</td>
<td>提现ID</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/dw/withdraw/api/create*/
{
  &#34;address&#34;: &#34;0xde709f2102306220921060314715629080e2fb77&#34;,
  &#34;amount&#34;: &#34;0.05&#34;,
  &#34;currency&#34;: &#34;eth&#34;,
  &#34;fee&#34;: &#34;0.01&#34;
}
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 700
}
</code></pre>
<h4>
<a id="user-content-post-v1dwwithdraw-virtualwithdraw-idcancel-申请取消提现虚拟币" class="anchor" href="#post-v1dwwithdraw-virtualwithdraw-idcancel-%E7%94%B3%E8%AF%B7%E5%8F%96%E6%B6%88%E6%8F%90%E7%8E%B0%E8%99%9A%E6%8B%9F%E5%B8%81" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>POST /v1/dw/withdraw-virtual/{withdraw-id}/cancel 申请取消提现虚拟币</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>withdraw-id</td>
<td>true</td>
<td>long</td>
<td>提现ID，填在path中</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>data</td>
<td>false</td>
<td>long</td>
<td>提现 ID</td>
<td></td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* POST /v1/dw/withdraw-virtual/{withdraw-id}/cancel */
{
  &#34;status&#34;: &#34;ok&#34;,
  &#34;data&#34;: 700
}
</code></pre>
<h4>
<a id="user-content-get-v1querydeposit-withdraw-查询虚拟币充提记录" class="anchor" href="#get-v1querydeposit-withdraw-%E6%9F%A5%E8%AF%A2%E8%99%9A%E6%8B%9F%E5%B8%81%E5%85%85%E6%8F%90%E8%AE%B0%E5%BD%95" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>GET /v1/query/deposit-withdraw 查询虚拟币充提记录</h4>
<p>请求参数:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>类型</th>
<th>描述</th>
<th>默认值</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
<td></td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>string</td>
<td>&#39;deposit&#39; or &#39;withdraw&#39;</td>
<td></td>
<td></td>
</tr>
<tr>
<td>from</td>
<td>false</td>
<td>string</td>
<td>查询起始 ID</td>
<td></td>
<td></td>
</tr>
<tr>
<td>size</td>
<td>false</td>
<td>string</td>
<td>查询记录大小</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<p>响应数据:</p>
<table>
<thead>
<tr>
<th>参数名称</th>
<th>是否必须</th>
<th>数据类型</th>
<th>描述</th>
<th>取值范围</th>
</tr>
</thead>
<tbody>
<tr>
<td>id</td>
<td>true</td>
<td>long</td>
<td></td>
<td></td>
</tr>
<tr>
<td>type</td>
<td>true</td>
<td>long</td>
<td>类型</td>
<td>&#39;deposit&#39; &#39;withdraw&#39;</td>
</tr>
<tr>
<td>currency</td>
<td>true</td>
<td>string</td>
<td>币种</td>
<td></td>
</tr>
<tr>
<td>tx-hash</td>
<td>true</td>
<td>string</td>
<td>交易哈希</td>
<td></td>
</tr>
<tr>
<td>amount</td>
<td>true</td>
<td>long</td>
<td>个数</td>
<td></td>
</tr>
<tr>
<td>address</td>
<td>true</td>
<td>string</td>
<td>地址</td>
<td></td>
</tr>
<tr>
<td>address-tag</td>
<td>true</td>
<td>string</td>
<td>地址标签</td>
<td></td>
</tr>
<tr>
<td>fee</td>
<td>true</td>
<td>long</td>
<td>手续费</td>
<td></td>
</tr>
<tr>
<td>state</td>
<td>true</td>
<td>string</td>
<td>状态</td>
<td>状态参见下表</td>
</tr>
<tr>
<td>created-at</td>
<td>true</td>
<td>long</td>
<td>发起时间</td>
<td></td>
</tr>
<tr>
<td>updated-at</td>
<td>true</td>
<td>long</td>
<td>最后更新时间</td>
<td></td>
</tr>
</tbody>
</table>
<h6>
<a id="user-content-虚拟币提现状态定义" class="anchor" href="#%E8%99%9A%E6%8B%9F%E5%B8%81%E6%8F%90%E7%8E%B0%E7%8A%B6%E6%80%81%E5%AE%9A%E4%B9%89" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>虚拟币提现状态定义：</h6>
<table>
<thead>
<tr>
<th>状态</th>
<th>描述</th>
</tr>
</thead>
<tbody>
<tr>
<td>submitted</td>
<td>已提交</td>
</tr>
<tr>
<td>reexamine</td>
<td>审核中</td>
</tr>
<tr>
<td>canceled</td>
<td>已撤销</td>
</tr>
<tr>
<td>pass</td>
<td>审批通过</td>
</tr>
<tr>
<td>reject</td>
<td>审批拒绝</td>
</tr>
<tr>
<td>pre-transfer</td>
<td>处理中</td>
</tr>
<tr>
<td>wallet-transfer</td>
<td>已汇出</td>
</tr>
<tr>
<td>wallet-reject</td>
<td>钱包拒绝</td>
</tr>
<tr>
<td>confirmed</td>
<td>区块已确认</td>
</tr>
<tr>
<td>confirm-error</td>
<td>区块确认错误</td>
</tr>
<tr>
<td>repealed</td>
<td>已撤销</td>
</tr>
</tbody>
</table>
<h6>
<a id="user-content-虚拟币充值状态定义" class="anchor" href="#%E8%99%9A%E6%8B%9F%E5%B8%81%E5%85%85%E5%80%BC%E7%8A%B6%E6%80%81%E5%AE%9A%E4%B9%89" aria-hidden="true"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>虚拟币充值状态定义：</h6>
<table>
<thead>
<tr>
<th>状态</th>
<th>描述</th>
</tr>
</thead>
<tbody>
<tr>
<td>unknown</td>
<td>状态未知</td>
</tr>
<tr>
<td>confirming</td>
<td>确认中</td>
</tr>
<tr>
<td>confirmed</td>
<td>确认中</td>
</tr>
<tr>
<td>safe</td>
<td>已完成</td>
</tr>
<tr>
<td>orphan</td>
<td>待确认</td>
</tr>
</tbody>
</table>
<p>请求响应例子:</p>
<pre><code>/* GET /v1/query/deposit-withdraw?currency=xrp&amp;type=deposit&amp;from=5&amp;size=12 */

{
    
    &#34;status&#34;: &#34;ok&#34;,
    &#34;data&#34;: [
      {
        &#34;id&#34;: 1171,
        &#34;type&#34;: &#34;deposit&#34;,
        &#34;currency&#34;: &#34;xrp&#34;,
        &#34;tx-hash&#34;: &#34;ed03094b84eafbe4bc16e7ef766ee959885ee5bcb265872baaa9c64e1cf86c2b&#34;,
        &#34;amount&#34;: 7.457467,
        &#34;address&#34;: &#34;rae93V8d2mdoUQHwBDBdM4NHCMehRJAsbm&#34;,
        &#34;address-tag&#34;: &#34;100040&#34;,
        &#34;fee&#34;: 0,
        &#34;state&#34;: &#34;safe&#34;,
        &#34;created-at&#34;: 1510912472199,
        &#34;updated-at&#34;: 1511145876575
      },
     ...
    ]
}
</code></pre>

        