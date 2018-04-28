<div id="itemContent" class="l-c-t">
                <div id="examples" class="item">
                    <h4>1.API 示例程序 <i class="caret"></i></h4>
                </div>
                <div>
                    <div class="sectioncont">
                        <p> PHP示例程序: <a target="_blank" href="https://github.com/gateio/rest/tree/master/php">https://github.com/gateio/rest/tree/master/php</a>
                        </p><p> JAVA示例程序: <a target="_blank" href="https://github.com/gateio/rest/tree/master/java">https://github.com/gateio/rest/tree/master/java</a>
                        </p><p> NODEJS示例程序: <a target="_blank" href="https://github.com/gateio/rest/tree/master/nodejs">https://github.com/gateio/rest/tree/master/nodejs</a>
                        </p><p> PYTHON示例程序: <a target="_blank" href="https://github.com/gateio/rest/tree/master/python">https://github.com/gateio/rest/tree/master/python</a>
                        </p><p> GO示例程序: <a target="_blank" href="https://github.com/gateio/rest/tree/master/go">https://github.com/gateio/rest/tree/master/go</a>
                        </p>
                    </div>
                </div>
                <div id="pairs" class="item">
                    <h4>2. 所有交易对 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回所有系统支持的交易对 </p>
                <p> URL: <a target="_blank" href="http://data.gateio.io/api2/1/pairs">http://data.gateio.io/api2/1/pairs</a>
                </p>
            </div>
                <h3>示例</h3>
            <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/pairs
    <span style="color:#0088cc"># Response</span>
        [
                "eth_btc","etc_btc","etc_eth","zec_btc","dash_btc","ltc_btc","bcc_btc","qtum_btc",
                "qtum_eth","xrp_btc","zrx_btc","zrx_eth","dnt_eth","dpy_eth","oax_eth","lrc_eth",
                "lrc_btc","pst_eth","tnt_eth","snt_eth","snt_btc","omg_eth","omg_btc","pay_eth",
                "pay_btc","bat_eth","cvc_eth","storj_eth","storj_btc","eos_eth","eos_btc"
        ]
            </pre>
            <h3>返回值说明</h3>
    <pre>    eth_btc: 以太币对比特币交易
    etc_btc: 以太经典对比特币
    etc_eth: 以太经典对以太币
    xrp_btc: 瑞波币对比特币
    zec_btc: ZCash对比特币
    ……
    </pre>
                </div>
                <div id="marketinfo" class="item">
                    <h4>3. 交易市场订单参数 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回所有系统支持的交易市场的参数信息，包括交易费，最小下单量，价格精度等。</p>
                <p> URL: <a target="_blank" href="http://data.gateio.io/api2/1/marketinfo">http://data.gateio.io/api2/1/marketinfo</a>
                </p>
            </div>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/marketinfo
    <span style="color:#0088cc"># Response</span>
   {
        "result": "true",
        "pairs": [
                      {
                            "eth_btc": {
                                "decimal_places": 6,
                                "min_amount": 0.0001,
                                "fee": 0.2
                            }
                      },
                      {
                            "etc_btc": {
                                "decimal_places": 6,
                                "min_amount": 0.0001,
                                "fee": 0.2
                            }
                      }
            ]
    }
            </pre>
            <h3>返回值说明</h3>
            <pre>    decimal_places: 价格精度
    min_amount : 最小下单量
    fee : 交易费</pre>
            <br>
                </div>
                <div id="marketlist" class="item">
                    <h4>4. 交易市场详细行情 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回所有系统支持的交易市场的详细行情和币种信息，包括币种名，市值，供应量，最新价格，涨跌趋势，价格曲线等。</p>
                <p> URL: <a target="_blank" href="http://data.gateio.io/api2/1/marketlist">http://data.gateio.io/api2/1/marketlist</a>
                </p>
            </div>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/marketlist
    <span style="color:#0088cc"># Response</span>
       {
        "result": "true",
        "data": [
                {
                    "no": 1,
                    "symbol": "LTC",
                    "name": "Litecoin",
                    "name_en": "Litecoin",
                    "name_cn": "莱特币",
                    "pair": "ltc_btc",
                    "rate": "418.00",
                    "vol_a": 168120.2,
                    "vol_b": "65,616,561",
                    "curr_a": "LTC",
                    "curr_b": "BTC",
                    "curr_suffix": " BTC",
                    "rate_percent": "19.73",
                    "trend": "up",
                    "supply": 25760300,
                    "marketcap": "10,767,805,404",
                    "plot": null
                },
                {
                    "no": 2,
                    "symbol": "ETH",
                    "name": "ETH",
                    "name_en": "ETH",
                    "name_cn": "以太币",
                    "pair": "etc_eth",
                    "rate": "0.7450",
                    "vol_a": 65227328.3,
                    "vol_b": "51,041,999",
                    "curr_a": "etc",
                    "curr_b": "eth",
                    "curr_suffix": " eth",
                    "rate_percent": "-1.84",
                    "trend": "up",
                    "supply": 1050000000,
                    "marketcap": "782,250,000",
                    "plot": null
                }
            ]
    }
            </pre>
            <h3>返回值说明</h3>
            <pre>    symbol : 币种标识
    name: 币种名称
    name_en: 英文名称
    name_cn: 中文名称
    pair: 交易对
    rate: 当前价格
    vol_a: 被兑换货币交易量
    vol_b: 兑换货币交易量
    curr_a: 被兑换货币
    curr_b: 兑换货币
    curr_suffix: 货币类型后缀
    rate_percent: 涨跌百分百
    trend: 24小时趋势 up涨 down跌
    supply: 币种供应量
    marketcap: 总市值
    plot: 趋势数据</pre>
            <br>
                </div>
                <div id="tickers" class="item">
                    <h4>5. 所有交易行情 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回系统支持的所有交易对的 最新，最高，最低 交易行情和交易量，每10秒钟更新: </p>
                <p> URL: <a target="_blank" href="http://data.gateio.io/api2/1/tickers">http://data.gateio.io/api2/1/tickers</a></p>
            </div>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/tickers
    <span style="color:#0088cc"># Response</span>
      {
          "eth_btc": {
                "result": "true",
                "last": 0.1,
                "lowestAsk": 0.1,
                "highestBid": "0.00000000",
                "percentChange": 0,
                "baseVolume": 0.001,
                "quoteVolume": 0.01,
                "high24hr": 0.1,
                "low24hr": 0.1
            },
        "xrp_btc": {
                "result": "true",
                "last": "0.00004720",
                "lowestAsk": "0.00005620",
                "highestBid": "0.00004550",
                "percentChange": -7.4510907045863,
                "baseVolume": 0.5324,
                "quoteVolume": 11417.333,
                "high24hr": "0.00005580",
                "low24hr": "0.00004560"
            }
        }
            </pre>
                <h3>返回值说明</h3>
            <pre>    baseVolume: 交易量
    high24hr:24小时最高价
    highestBid:买方最高价
    last:最新成交价
    low24hr:24小时最低价
    lowestAsk:卖方最低价
    percentChange:涨跌百分比
    quoteVolume: 兑换货币交易量</pre>
            <br>
            </div>

                <div id="ticker" class="item">
                    <h4>6. 单项交易行情 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回最新，最高，最低 交易行情和交易量，每10秒钟更新: </p>
                <p> URL: http://data.gateio.io/api2/1/ticker/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]</font></p>
                <p> 请替换 [CURR_A] and [CURR_B] 为您需要查看的币种.</p>
                <p> 支持的兑换类型: </p>
                <ol class="type-ol">
					<li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/usdt_cny">http://data.gateio.io/api2/1/ticker/<span>usdt_cny</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btc_usdt">http://data.gateio.io/api2/1/ticker/<span>btc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bch_usdt">http://data.gateio.io/api2/1/ticker/<span>bch_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/eth_usdt">http://data.gateio.io/api2/1/ticker/<span>eth_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/etc_usdt">http://data.gateio.io/api2/1/ticker/<span>etc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qtum_usdt">http://data.gateio.io/api2/1/ticker/<span>qtum_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ltc_usdt">http://data.gateio.io/api2/1/ticker/<span>ltc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dash_usdt">http://data.gateio.io/api2/1/ticker/<span>dash_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zec_usdt">http://data.gateio.io/api2/1/ticker/<span>zec_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btm_usdt">http://data.gateio.io/api2/1/ticker/<span>btm_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/eos_usdt">http://data.gateio.io/api2/1/ticker/<span>eos_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/req_usdt">http://data.gateio.io/api2/1/ticker/<span>req_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/snt_usdt">http://data.gateio.io/api2/1/ticker/<span>snt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/omg_usdt">http://data.gateio.io/api2/1/ticker/<span>omg_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/pay_usdt">http://data.gateio.io/api2/1/ticker/<span>pay_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cvc_usdt">http://data.gateio.io/api2/1/ticker/<span>cvc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zrx_usdt">http://data.gateio.io/api2/1/ticker/<span>zrx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tnt_usdt">http://data.gateio.io/api2/1/ticker/<span>tnt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xmr_usdt">http://data.gateio.io/api2/1/ticker/<span>xmr_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xrp_usdt">http://data.gateio.io/api2/1/ticker/<span>xrp_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/doge_usdt">http://data.gateio.io/api2/1/ticker/<span>doge_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bat_usdt">http://data.gateio.io/api2/1/ticker/<span>bat_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/pst_usdt">http://data.gateio.io/api2/1/ticker/<span>pst_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btg_usdt">http://data.gateio.io/api2/1/ticker/<span>btg_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dpy_usdt">http://data.gateio.io/api2/1/ticker/<span>dpy_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lrc_usdt">http://data.gateio.io/api2/1/ticker/<span>lrc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/storj_usdt">http://data.gateio.io/api2/1/ticker/<span>storj_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rdn_usdt">http://data.gateio.io/api2/1/ticker/<span>rdn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/stx_usdt">http://data.gateio.io/api2/1/ticker/<span>stx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/knc_usdt">http://data.gateio.io/api2/1/ticker/<span>knc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/link_usdt">http://data.gateio.io/api2/1/ticker/<span>link_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cdt_usdt">http://data.gateio.io/api2/1/ticker/<span>cdt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ae_usdt">http://data.gateio.io/api2/1/ticker/<span>ae_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ae_eth">http://data.gateio.io/api2/1/ticker/<span>ae_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ae_btc">http://data.gateio.io/api2/1/ticker/<span>ae_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cdt_eth">http://data.gateio.io/api2/1/ticker/<span>cdt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rdn_eth">http://data.gateio.io/api2/1/ticker/<span>rdn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/stx_eth">http://data.gateio.io/api2/1/ticker/<span>stx_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/knc_eth">http://data.gateio.io/api2/1/ticker/<span>knc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/link_eth">http://data.gateio.io/api2/1/ticker/<span>link_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/req_eth">http://data.gateio.io/api2/1/ticker/<span>req_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rcn_eth">http://data.gateio.io/api2/1/ticker/<span>rcn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/trx_eth">http://data.gateio.io/api2/1/ticker/<span>trx_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/arn_eth">http://data.gateio.io/api2/1/ticker/<span>arn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/kick_eth">http://data.gateio.io/api2/1/ticker/<span>kick_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bnt_eth">http://data.gateio.io/api2/1/ticker/<span>bnt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ven_eth">http://data.gateio.io/api2/1/ticker/<span>ven_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mco_eth">http://data.gateio.io/api2/1/ticker/<span>mco_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/fun_eth">http://data.gateio.io/api2/1/ticker/<span>fun_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/data_eth">http://data.gateio.io/api2/1/ticker/<span>data_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rlc_eth">http://data.gateio.io/api2/1/ticker/<span>rlc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zsc_eth">http://data.gateio.io/api2/1/ticker/<span>zsc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/wings_eth">http://data.gateio.io/api2/1/ticker/<span>wings_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ctr_eth">http://data.gateio.io/api2/1/ticker/<span>ctr_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mda_eth">http://data.gateio.io/api2/1/ticker/<span>mda_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rcn_usdt">http://data.gateio.io/api2/1/ticker/<span>rcn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/trx_usdt">http://data.gateio.io/api2/1/ticker/<span>trx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/kick_usdt">http://data.gateio.io/api2/1/ticker/<span>kick_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ven_usdt">http://data.gateio.io/api2/1/ticker/<span>ven_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mco_usdt">http://data.gateio.io/api2/1/ticker/<span>mco_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/fun_usdt">http://data.gateio.io/api2/1/ticker/<span>fun_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/data_usdt">http://data.gateio.io/api2/1/ticker/<span>data_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zsc_usdt">http://data.gateio.io/api2/1/ticker/<span>zsc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mda_usdt">http://data.gateio.io/api2/1/ticker/<span>mda_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xtz_usdt">http://data.gateio.io/api2/1/ticker/<span>xtz_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gnt_usdt">http://data.gateio.io/api2/1/ticker/<span>gnt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gnt_eth">http://data.gateio.io/api2/1/ticker/<span>gnt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gem_usdt">http://data.gateio.io/api2/1/ticker/<span>gem_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gem_eth">http://data.gateio.io/api2/1/ticker/<span>gem_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rfr_usdt">http://data.gateio.io/api2/1/ticker/<span>rfr_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rfr_eth">http://data.gateio.io/api2/1/ticker/<span>rfr_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dadi_usdt">http://data.gateio.io/api2/1/ticker/<span>dadi_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dadi_eth">http://data.gateio.io/api2/1/ticker/<span>dadi_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/abt_usdt">http://data.gateio.io/api2/1/ticker/<span>abt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/abt_eth">http://data.gateio.io/api2/1/ticker/<span>abt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ledu_usdt">http://data.gateio.io/api2/1/ticker/<span>ledu_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ledu_btc">http://data.gateio.io/api2/1/ticker/<span>ledu_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ledu_eth">http://data.gateio.io/api2/1/ticker/<span>ledu_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ost_usdt">http://data.gateio.io/api2/1/ticker/<span>ost_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ost_eth">http://data.gateio.io/api2/1/ticker/<span>ost_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xlm_usdt">http://data.gateio.io/api2/1/ticker/<span>xlm_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xlm_eth">http://data.gateio.io/api2/1/ticker/<span>xlm_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xlm_btc">http://data.gateio.io/api2/1/ticker/<span>xlm_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mobi_usdt">http://data.gateio.io/api2/1/ticker/<span>mobi_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mobi_eth">http://data.gateio.io/api2/1/ticker/<span>mobi_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mobi_btc">http://data.gateio.io/api2/1/ticker/<span>mobi_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ocn_usdt">http://data.gateio.io/api2/1/ticker/<span>ocn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ocn_eth">http://data.gateio.io/api2/1/ticker/<span>ocn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ocn_btc">http://data.gateio.io/api2/1/ticker/<span>ocn_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zpt_usdt">http://data.gateio.io/api2/1/ticker/<span>zpt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zpt_eth">http://data.gateio.io/api2/1/ticker/<span>zpt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zpt_btc">http://data.gateio.io/api2/1/ticker/<span>zpt_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cofi_usdt">http://data.gateio.io/api2/1/ticker/<span>cofi_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cofi_eth">http://data.gateio.io/api2/1/ticker/<span>cofi_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/jnt_usdt">http://data.gateio.io/api2/1/ticker/<span>jnt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/jnt_eth">http://data.gateio.io/api2/1/ticker/<span>jnt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/jnt_btc">http://data.gateio.io/api2/1/ticker/<span>jnt_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/blz_usdt">http://data.gateio.io/api2/1/ticker/<span>blz_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/blz_eth">http://data.gateio.io/api2/1/ticker/<span>blz_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gxs_usdt">http://data.gateio.io/api2/1/ticker/<span>gxs_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gxs_btc">http://data.gateio.io/api2/1/ticker/<span>gxs_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mtn_usdt">http://data.gateio.io/api2/1/ticker/<span>mtn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mtn_eth">http://data.gateio.io/api2/1/ticker/<span>mtn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ruff_usdt">http://data.gateio.io/api2/1/ticker/<span>ruff_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ruff_eth">http://data.gateio.io/api2/1/ticker/<span>ruff_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ruff_btc">http://data.gateio.io/api2/1/ticker/<span>ruff_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tnc_usdt">http://data.gateio.io/api2/1/ticker/<span>tnc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tnc_eth">http://data.gateio.io/api2/1/ticker/<span>tnc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tnc_btc">http://data.gateio.io/api2/1/ticker/<span>tnc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zil_usdt">http://data.gateio.io/api2/1/ticker/<span>zil_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zil_eth">http://data.gateio.io/api2/1/ticker/<span>zil_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tio_usdt">http://data.gateio.io/api2/1/ticker/<span>tio_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tio_eth">http://data.gateio.io/api2/1/ticker/<span>tio_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bto_usdt">http://data.gateio.io/api2/1/ticker/<span>bto_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bto_eth">http://data.gateio.io/api2/1/ticker/<span>bto_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/theta_usdt">http://data.gateio.io/api2/1/ticker/<span>theta_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/theta_eth">http://data.gateio.io/api2/1/ticker/<span>theta_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ddd_usdt">http://data.gateio.io/api2/1/ticker/<span>ddd_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ddd_eth">http://data.gateio.io/api2/1/ticker/<span>ddd_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ddd_btc">http://data.gateio.io/api2/1/ticker/<span>ddd_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mkr_usdt">http://data.gateio.io/api2/1/ticker/<span>mkr_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mkr_eth">http://data.gateio.io/api2/1/ticker/<span>mkr_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dai_usdt">http://data.gateio.io/api2/1/ticker/<span>dai_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/smt_usdt">http://data.gateio.io/api2/1/ticker/<span>smt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/smt_eth">http://data.gateio.io/api2/1/ticker/<span>smt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mdt_usdt">http://data.gateio.io/api2/1/ticker/<span>mdt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mdt_eth">http://data.gateio.io/api2/1/ticker/<span>mdt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mdt_btc">http://data.gateio.io/api2/1/ticker/<span>mdt_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mana_usdt">http://data.gateio.io/api2/1/ticker/<span>mana_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mana_eth">http://data.gateio.io/api2/1/ticker/<span>mana_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lun_usdt">http://data.gateio.io/api2/1/ticker/<span>lun_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lun_eth">http://data.gateio.io/api2/1/ticker/<span>lun_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/salt_usdt">http://data.gateio.io/api2/1/ticker/<span>salt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/salt_eth">http://data.gateio.io/api2/1/ticker/<span>salt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/fuel_usdt">http://data.gateio.io/api2/1/ticker/<span>fuel_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/fuel_eth">http://data.gateio.io/api2/1/ticker/<span>fuel_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/elf_usdt">http://data.gateio.io/api2/1/ticker/<span>elf_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/elf_eth">http://data.gateio.io/api2/1/ticker/<span>elf_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/drgn_usdt">http://data.gateio.io/api2/1/ticker/<span>drgn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/drgn_eth">http://data.gateio.io/api2/1/ticker/<span>drgn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gtc_usdt">http://data.gateio.io/api2/1/ticker/<span>gtc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gtc_eth">http://data.gateio.io/api2/1/ticker/<span>gtc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gtc_btc">http://data.gateio.io/api2/1/ticker/<span>gtc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qlc_usdt">http://data.gateio.io/api2/1/ticker/<span>qlc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qlc_btc">http://data.gateio.io/api2/1/ticker/<span>qlc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qlc_eth">http://data.gateio.io/api2/1/ticker/<span>qlc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dbc_usdt">http://data.gateio.io/api2/1/ticker/<span>dbc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dbc_btc">http://data.gateio.io/api2/1/ticker/<span>dbc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dbc_eth">http://data.gateio.io/api2/1/ticker/<span>dbc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bnty_usdt">http://data.gateio.io/api2/1/ticker/<span>bnty_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bnty_eth">http://data.gateio.io/api2/1/ticker/<span>bnty_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lend_usdt">http://data.gateio.io/api2/1/ticker/<span>lend_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lend_eth">http://data.gateio.io/api2/1/ticker/<span>lend_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/icx_usdt">http://data.gateio.io/api2/1/ticker/<span>icx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/icx_eth">http://data.gateio.io/api2/1/ticker/<span>icx_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btf_usdt">http://data.gateio.io/api2/1/ticker/<span>btf_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btf_btc">http://data.gateio.io/api2/1/ticker/<span>btf_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ada_usdt">http://data.gateio.io/api2/1/ticker/<span>ada_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ada_btc">http://data.gateio.io/api2/1/ticker/<span>ada_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lsk_usdt">http://data.gateio.io/api2/1/ticker/<span>lsk_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lsk_btc">http://data.gateio.io/api2/1/ticker/<span>lsk_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/waves_usdt">http://data.gateio.io/api2/1/ticker/<span>waves_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/waves_btc">http://data.gateio.io/api2/1/ticker/<span>waves_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bifi_usdt">http://data.gateio.io/api2/1/ticker/<span>bifi_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bifi_btc">http://data.gateio.io/api2/1/ticker/<span>bifi_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mds_eth">http://data.gateio.io/api2/1/ticker/<span>mds_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/mds_usdt">http://data.gateio.io/api2/1/ticker/<span>mds_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dgd_usdt">http://data.gateio.io/api2/1/ticker/<span>dgd_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dgd_eth">http://data.gateio.io/api2/1/ticker/<span>dgd_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qash_usdt">http://data.gateio.io/api2/1/ticker/<span>qash_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qash_eth">http://data.gateio.io/api2/1/ticker/<span>qash_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qash_btc">http://data.gateio.io/api2/1/ticker/<span>qash_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/powr_usdt">http://data.gateio.io/api2/1/ticker/<span>powr_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/powr_eth">http://data.gateio.io/api2/1/ticker/<span>powr_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/powr_btc">http://data.gateio.io/api2/1/ticker/<span>powr_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/fil_usdt">http://data.gateio.io/api2/1/ticker/<span>fil_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcd_usdt">http://data.gateio.io/api2/1/ticker/<span>bcd_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcd_btc">http://data.gateio.io/api2/1/ticker/<span>bcd_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/sbtc_usdt">http://data.gateio.io/api2/1/ticker/<span>sbtc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/sbtc_btc">http://data.gateio.io/api2/1/ticker/<span>sbtc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/god_usdt">http://data.gateio.io/api2/1/ticker/<span>god_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/god_btc">http://data.gateio.io/api2/1/ticker/<span>god_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcx_usdt">http://data.gateio.io/api2/1/ticker/<span>bcx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcx_btc">http://data.gateio.io/api2/1/ticker/<span>bcx_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/hsr_usdt">http://data.gateio.io/api2/1/ticker/<span>hsr_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/hsr_btc">http://data.gateio.io/api2/1/ticker/<span>hsr_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/hsr_eth">http://data.gateio.io/api2/1/ticker/<span>hsr_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qsp_usdt">http://data.gateio.io/api2/1/ticker/<span>qsp_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qsp_eth">http://data.gateio.io/api2/1/ticker/<span>qsp_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ink_btc">http://data.gateio.io/api2/1/ticker/<span>ink_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ink_usdt">http://data.gateio.io/api2/1/ticker/<span>ink_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ink_eth">http://data.gateio.io/api2/1/ticker/<span>ink_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ink_qtum">http://data.gateio.io/api2/1/ticker/<span>ink_qtum</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/med_qtum">http://data.gateio.io/api2/1/ticker/<span>med_qtum</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/med_eth">http://data.gateio.io/api2/1/ticker/<span>med_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/med_usdt">http://data.gateio.io/api2/1/ticker/<span>med_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bot_qtum">http://data.gateio.io/api2/1/ticker/<span>bot_qtum</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bot_usdt">http://data.gateio.io/api2/1/ticker/<span>bot_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bot_eth">http://data.gateio.io/api2/1/ticker/<span>bot_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qbt_qtum">http://data.gateio.io/api2/1/ticker/<span>qbt_qtum</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qbt_eth">http://data.gateio.io/api2/1/ticker/<span>qbt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qbt_usdt">http://data.gateio.io/api2/1/ticker/<span>qbt_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tsl_qtum">http://data.gateio.io/api2/1/ticker/<span>tsl_qtum</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tsl_usdt">http://data.gateio.io/api2/1/ticker/<span>tsl_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gnx_usdt">http://data.gateio.io/api2/1/ticker/<span>gnx_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gnx_eth">http://data.gateio.io/api2/1/ticker/<span>gnx_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/neo_usdt">http://data.gateio.io/api2/1/ticker/<span>neo_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gas_usdt">http://data.gateio.io/api2/1/ticker/<span>gas_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/neo_btc">http://data.gateio.io/api2/1/ticker/<span>neo_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/gas_btc">http://data.gateio.io/api2/1/ticker/<span>gas_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/iota_usdt">http://data.gateio.io/api2/1/ticker/<span>iota_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/iota_btc">http://data.gateio.io/api2/1/ticker/<span>iota_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/nas_usdt">http://data.gateio.io/api2/1/ticker/<span>nas_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/nas_eth">http://data.gateio.io/api2/1/ticker/<span>nas_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/nas_btc">http://data.gateio.io/api2/1/ticker/<span>nas_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/eth_btc">http://data.gateio.io/api2/1/ticker/<span>eth_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/etc_btc">http://data.gateio.io/api2/1/ticker/<span>etc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/etc_eth">http://data.gateio.io/api2/1/ticker/<span>etc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zec_btc">http://data.gateio.io/api2/1/ticker/<span>zec_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dash_btc">http://data.gateio.io/api2/1/ticker/<span>dash_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ltc_btc">http://data.gateio.io/api2/1/ticker/<span>ltc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bch_btc">http://data.gateio.io/api2/1/ticker/<span>bch_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btg_btc">http://data.gateio.io/api2/1/ticker/<span>btg_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qtum_btc">http://data.gateio.io/api2/1/ticker/<span>qtum_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/qtum_eth">http://data.gateio.io/api2/1/ticker/<span>qtum_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xrp_btc">http://data.gateio.io/api2/1/ticker/<span>xrp_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/doge_btc">http://data.gateio.io/api2/1/ticker/<span>doge_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xmr_btc">http://data.gateio.io/api2/1/ticker/<span>xmr_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zrx_btc">http://data.gateio.io/api2/1/ticker/<span>zrx_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/zrx_eth">http://data.gateio.io/api2/1/ticker/<span>zrx_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dnt_eth">http://data.gateio.io/api2/1/ticker/<span>dnt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dpy_eth">http://data.gateio.io/api2/1/ticker/<span>dpy_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/oax_eth">http://data.gateio.io/api2/1/ticker/<span>oax_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rep_eth">http://data.gateio.io/api2/1/ticker/<span>rep_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lrc_eth">http://data.gateio.io/api2/1/ticker/<span>lrc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lrc_btc">http://data.gateio.io/api2/1/ticker/<span>lrc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/pst_eth">http://data.gateio.io/api2/1/ticker/<span>pst_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcdn_eth">http://data.gateio.io/api2/1/ticker/<span>bcdn_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bcdn_usdt">http://data.gateio.io/api2/1/ticker/<span>bcdn_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tnt_eth">http://data.gateio.io/api2/1/ticker/<span>tnt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/snt_eth">http://data.gateio.io/api2/1/ticker/<span>snt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/snt_btc">http://data.gateio.io/api2/1/ticker/<span>snt_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btm_eth">http://data.gateio.io/api2/1/ticker/<span>btm_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/btm_btc">http://data.gateio.io/api2/1/ticker/<span>btm_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/llt_eth">http://data.gateio.io/api2/1/ticker/<span>llt_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/snet_eth">http://data.gateio.io/api2/1/ticker/<span>snet_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/snet_usdt">http://data.gateio.io/api2/1/ticker/<span>snet_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/llt_snet">http://data.gateio.io/api2/1/ticker/<span>llt_snet</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/omg_eth">http://data.gateio.io/api2/1/ticker/<span>omg_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/omg_btc">http://data.gateio.io/api2/1/ticker/<span>omg_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/pay_eth">http://data.gateio.io/api2/1/ticker/<span>pay_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/pay_btc">http://data.gateio.io/api2/1/ticker/<span>pay_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bat_eth">http://data.gateio.io/api2/1/ticker/<span>bat_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bat_btc">http://data.gateio.io/api2/1/ticker/<span>bat_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cvc_eth">http://data.gateio.io/api2/1/ticker/<span>cvc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/storj_eth">http://data.gateio.io/api2/1/ticker/<span>storj_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/storj_btc">http://data.gateio.io/api2/1/ticker/<span>storj_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/eos_eth">http://data.gateio.io/api2/1/ticker/<span>eos_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/eos_btc">http://data.gateio.io/api2/1/ticker/<span>eos_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bts_usdt">http://data.gateio.io/api2/1/ticker/<span>bts_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bts_btc">http://data.gateio.io/api2/1/ticker/<span>bts_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tips_eth">http://data.gateio.io/api2/1/ticker/<span>tips_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xmc_usdt">http://data.gateio.io/api2/1/ticker/<span>xmc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/xmc_btc">http://data.gateio.io/api2/1/ticker/<span>xmc_btc</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cs_eth">http://data.gateio.io/api2/1/ticker/<span>cs_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/cs_usdt">http://data.gateio.io/api2/1/ticker/<span>cs_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/man_eth">http://data.gateio.io/api2/1/ticker/<span>man_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/man_usdt">http://data.gateio.io/api2/1/ticker/<span>man_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/rem_eth">http://data.gateio.io/api2/1/ticker/<span>rem_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lym_eth">http://data.gateio.io/api2/1/ticker/<span>lym_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lym_usdt">http://data.gateio.io/api2/1/ticker/<span>lym_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/instar_eth">http://data.gateio.io/api2/1/ticker/<span>instar_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ont_eth">http://data.gateio.io/api2/1/ticker/<span>ont_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ont_usdt">http://data.gateio.io/api2/1/ticker/<span>ont_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bft_eth">http://data.gateio.io/api2/1/ticker/<span>bft_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/bft_usdt">http://data.gateio.io/api2/1/ticker/<span>bft_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/iht_eth">http://data.gateio.io/api2/1/ticker/<span>iht_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/iht_usdt">http://data.gateio.io/api2/1/ticker/<span>iht_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/senc_eth">http://data.gateio.io/api2/1/ticker/<span>senc_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/senc_usdt">http://data.gateio.io/api2/1/ticker/<span>senc_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tomo_eth">http://data.gateio.io/api2/1/ticker/<span>tomo_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/elec_eth">http://data.gateio.io/api2/1/ticker/<span>elec_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/ship_eth">http://data.gateio.io/api2/1/ticker/<span>ship_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/tfd_eth">http://data.gateio.io/api2/1/ticker/<span>tfd_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/hav_eth">http://data.gateio.io/api2/1/ticker/<span>hav_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/hur_eth">http://data.gateio.io/api2/1/ticker/<span>hur_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lst_eth">http://data.gateio.io/api2/1/ticker/<span>lst_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/lino_eth">http://data.gateio.io/api2/1/ticker/<span>lino_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/swh_eth">http://data.gateio.io/api2/1/ticker/<span>swh_eth</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/swh_usdt">http://data.gateio.io/api2/1/ticker/<span>swh_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dock_usdt">http://data.gateio.io/api2/1/ticker/<span>dock_usdt</span></a></li>
                					
                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/ticker/dock_eth">http://data.gateio.io/api2/1/ticker/<span>dock_eth</span></a></li>
                					
                </ol>
            </div>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/ticker/eth_btc
    <span style="color:#0088cc"># Response</span>
        {
            "result": "true",
            "last": 0.1,
            "lowestAsk": 0.1,
            "highestBid": "0.00000000",
            "percentChange": 0,
            "baseVolume": 0.001,
            "quoteVolume": 0.01,
            "high24hr": 0.1,
            "low24hr": 0.1
        }
            </pre>
                <h3>返回值说明</h3>
                <pre>    baseVolume: 交易量
    high24hr:24小时最高价
    highestBid:买方最高价
    last:最新成交价
    low24hr:24小时最低价
    lowestAsk:卖方最低价
    percentChange:涨跌百分比
    quoteVolume: 兑换货币交易量</pre>
            <br>
                </div>
                <div id="depth" class="item">
                    <h4>7. 市场深度 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回系统支持的所有交易对的市场深度（委托挂单），其中 asks 是委卖单, bids 是委买单。 </p>
                <p> URL: <a target="_blank" href="http://data.gateio.io/api2/1/orderBooks">http://data.gateio.io/api2/1/orderBooks</a>
                </p>
            </div>
            <br>

            <div id="sidebartitlewrapper"><h3>当前市场深度 API</h3></div>
            <div class="sectioncont">
                <p> 返回当前市场深度（委托挂单），其中 asks 是委卖单, bids 是委买单。 </p>
                <p> URL: http://data.gateio.io/api2/1/orderBook/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]</font></p>
                <p> 请替换 [CURR_A] and [CURR_B] 为您需要查看的币种.</p>
                <p> 支持的兑换类型: </p>
                <ol class="type-ol">
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btc_usdt">http://data.gateio.io/api2/1/orderBook/<span>btc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bch_usdt">http://data.gateio.io/api2/1/orderBook/<span>bch_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/eth_usdt">http://data.gateio.io/api2/1/orderBook/<span>eth_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/etc_usdt">http://data.gateio.io/api2/1/orderBook/<span>etc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qtum_usdt">http://data.gateio.io/api2/1/orderBook/<span>qtum_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ltc_usdt">http://data.gateio.io/api2/1/orderBook/<span>ltc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dash_usdt">http://data.gateio.io/api2/1/orderBook/<span>dash_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zec_usdt">http://data.gateio.io/api2/1/orderBook/<span>zec_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btm_usdt">http://data.gateio.io/api2/1/orderBook/<span>btm_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/eos_usdt">http://data.gateio.io/api2/1/orderBook/<span>eos_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/req_usdt">http://data.gateio.io/api2/1/orderBook/<span>req_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/snt_usdt">http://data.gateio.io/api2/1/orderBook/<span>snt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/omg_usdt">http://data.gateio.io/api2/1/orderBook/<span>omg_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/pay_usdt">http://data.gateio.io/api2/1/orderBook/<span>pay_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cvc_usdt">http://data.gateio.io/api2/1/orderBook/<span>cvc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zrx_usdt">http://data.gateio.io/api2/1/orderBook/<span>zrx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tnt_usdt">http://data.gateio.io/api2/1/orderBook/<span>tnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xmr_usdt">http://data.gateio.io/api2/1/orderBook/<span>xmr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xrp_usdt">http://data.gateio.io/api2/1/orderBook/<span>xrp_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/doge_usdt">http://data.gateio.io/api2/1/orderBook/<span>doge_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bat_usdt">http://data.gateio.io/api2/1/orderBook/<span>bat_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/pst_usdt">http://data.gateio.io/api2/1/orderBook/<span>pst_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btg_usdt">http://data.gateio.io/api2/1/orderBook/<span>btg_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dpy_usdt">http://data.gateio.io/api2/1/orderBook/<span>dpy_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lrc_usdt">http://data.gateio.io/api2/1/orderBook/<span>lrc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/storj_usdt">http://data.gateio.io/api2/1/orderBook/<span>storj_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rdn_usdt">http://data.gateio.io/api2/1/orderBook/<span>rdn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/stx_usdt">http://data.gateio.io/api2/1/orderBook/<span>stx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/knc_usdt">http://data.gateio.io/api2/1/orderBook/<span>knc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/link_usdt">http://data.gateio.io/api2/1/orderBook/<span>link_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cdt_usdt">http://data.gateio.io/api2/1/orderBook/<span>cdt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ae_usdt">http://data.gateio.io/api2/1/orderBook/<span>ae_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ae_eth">http://data.gateio.io/api2/1/orderBook/<span>ae_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ae_btc">http://data.gateio.io/api2/1/orderBook/<span>ae_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cdt_eth">http://data.gateio.io/api2/1/orderBook/<span>cdt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rdn_eth">http://data.gateio.io/api2/1/orderBook/<span>rdn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/stx_eth">http://data.gateio.io/api2/1/orderBook/<span>stx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/knc_eth">http://data.gateio.io/api2/1/orderBook/<span>knc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/link_eth">http://data.gateio.io/api2/1/orderBook/<span>link_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/req_eth">http://data.gateio.io/api2/1/orderBook/<span>req_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rcn_eth">http://data.gateio.io/api2/1/orderBook/<span>rcn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/trx_eth">http://data.gateio.io/api2/1/orderBook/<span>trx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/arn_eth">http://data.gateio.io/api2/1/orderBook/<span>arn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/kick_eth">http://data.gateio.io/api2/1/orderBook/<span>kick_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bnt_eth">http://data.gateio.io/api2/1/orderBook/<span>bnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ven_eth">http://data.gateio.io/api2/1/orderBook/<span>ven_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mco_eth">http://data.gateio.io/api2/1/orderBook/<span>mco_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/fun_eth">http://data.gateio.io/api2/1/orderBook/<span>fun_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/data_eth">http://data.gateio.io/api2/1/orderBook/<span>data_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rlc_eth">http://data.gateio.io/api2/1/orderBook/<span>rlc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zsc_eth">http://data.gateio.io/api2/1/orderBook/<span>zsc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/wings_eth">http://data.gateio.io/api2/1/orderBook/<span>wings_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ctr_eth">http://data.gateio.io/api2/1/orderBook/<span>ctr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mda_eth">http://data.gateio.io/api2/1/orderBook/<span>mda_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rcn_usdt">http://data.gateio.io/api2/1/orderBook/<span>rcn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/trx_usdt">http://data.gateio.io/api2/1/orderBook/<span>trx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/kick_usdt">http://data.gateio.io/api2/1/orderBook/<span>kick_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ven_usdt">http://data.gateio.io/api2/1/orderBook/<span>ven_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mco_usdt">http://data.gateio.io/api2/1/orderBook/<span>mco_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/fun_usdt">http://data.gateio.io/api2/1/orderBook/<span>fun_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/data_usdt">http://data.gateio.io/api2/1/orderBook/<span>data_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zsc_usdt">http://data.gateio.io/api2/1/orderBook/<span>zsc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mda_usdt">http://data.gateio.io/api2/1/orderBook/<span>mda_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xtz_usdt">http://data.gateio.io/api2/1/orderBook/<span>xtz_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gnt_usdt">http://data.gateio.io/api2/1/orderBook/<span>gnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gnt_eth">http://data.gateio.io/api2/1/orderBook/<span>gnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gem_usdt">http://data.gateio.io/api2/1/orderBook/<span>gem_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gem_eth">http://data.gateio.io/api2/1/orderBook/<span>gem_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rfr_usdt">http://data.gateio.io/api2/1/orderBook/<span>rfr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rfr_eth">http://data.gateio.io/api2/1/orderBook/<span>rfr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dadi_usdt">http://data.gateio.io/api2/1/orderBook/<span>dadi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dadi_eth">http://data.gateio.io/api2/1/orderBook/<span>dadi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/abt_usdt">http://data.gateio.io/api2/1/orderBook/<span>abt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/abt_eth">http://data.gateio.io/api2/1/orderBook/<span>abt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ledu_usdt">http://data.gateio.io/api2/1/orderBook/<span>ledu_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ledu_btc">http://data.gateio.io/api2/1/orderBook/<span>ledu_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ledu_eth">http://data.gateio.io/api2/1/orderBook/<span>ledu_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ost_usdt">http://data.gateio.io/api2/1/orderBook/<span>ost_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ost_eth">http://data.gateio.io/api2/1/orderBook/<span>ost_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xlm_usdt">http://data.gateio.io/api2/1/orderBook/<span>xlm_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xlm_eth">http://data.gateio.io/api2/1/orderBook/<span>xlm_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xlm_btc">http://data.gateio.io/api2/1/orderBook/<span>xlm_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mobi_usdt">http://data.gateio.io/api2/1/orderBook/<span>mobi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mobi_eth">http://data.gateio.io/api2/1/orderBook/<span>mobi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mobi_btc">http://data.gateio.io/api2/1/orderBook/<span>mobi_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ocn_usdt">http://data.gateio.io/api2/1/orderBook/<span>ocn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ocn_eth">http://data.gateio.io/api2/1/orderBook/<span>ocn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ocn_btc">http://data.gateio.io/api2/1/orderBook/<span>ocn_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zpt_usdt">http://data.gateio.io/api2/1/orderBook/<span>zpt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zpt_eth">http://data.gateio.io/api2/1/orderBook/<span>zpt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zpt_btc">http://data.gateio.io/api2/1/orderBook/<span>zpt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cofi_usdt">http://data.gateio.io/api2/1/orderBook/<span>cofi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cofi_eth">http://data.gateio.io/api2/1/orderBook/<span>cofi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/jnt_usdt">http://data.gateio.io/api2/1/orderBook/<span>jnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/jnt_eth">http://data.gateio.io/api2/1/orderBook/<span>jnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/jnt_btc">http://data.gateio.io/api2/1/orderBook/<span>jnt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/blz_usdt">http://data.gateio.io/api2/1/orderBook/<span>blz_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/blz_eth">http://data.gateio.io/api2/1/orderBook/<span>blz_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gxs_usdt">http://data.gateio.io/api2/1/orderBook/<span>gxs_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gxs_btc">http://data.gateio.io/api2/1/orderBook/<span>gxs_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mtn_usdt">http://data.gateio.io/api2/1/orderBook/<span>mtn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mtn_eth">http://data.gateio.io/api2/1/orderBook/<span>mtn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ruff_usdt">http://data.gateio.io/api2/1/orderBook/<span>ruff_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ruff_eth">http://data.gateio.io/api2/1/orderBook/<span>ruff_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ruff_btc">http://data.gateio.io/api2/1/orderBook/<span>ruff_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tnc_usdt">http://data.gateio.io/api2/1/orderBook/<span>tnc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tnc_eth">http://data.gateio.io/api2/1/orderBook/<span>tnc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tnc_btc">http://data.gateio.io/api2/1/orderBook/<span>tnc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zil_usdt">http://data.gateio.io/api2/1/orderBook/<span>zil_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zil_eth">http://data.gateio.io/api2/1/orderBook/<span>zil_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tio_usdt">http://data.gateio.io/api2/1/orderBook/<span>tio_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tio_eth">http://data.gateio.io/api2/1/orderBook/<span>tio_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bto_usdt">http://data.gateio.io/api2/1/orderBook/<span>bto_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bto_eth">http://data.gateio.io/api2/1/orderBook/<span>bto_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/theta_usdt">http://data.gateio.io/api2/1/orderBook/<span>theta_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/theta_eth">http://data.gateio.io/api2/1/orderBook/<span>theta_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ddd_usdt">http://data.gateio.io/api2/1/orderBook/<span>ddd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ddd_eth">http://data.gateio.io/api2/1/orderBook/<span>ddd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ddd_btc">http://data.gateio.io/api2/1/orderBook/<span>ddd_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mkr_usdt">http://data.gateio.io/api2/1/orderBook/<span>mkr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mkr_eth">http://data.gateio.io/api2/1/orderBook/<span>mkr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dai_usdt">http://data.gateio.io/api2/1/orderBook/<span>dai_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/smt_usdt">http://data.gateio.io/api2/1/orderBook/<span>smt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/smt_eth">http://data.gateio.io/api2/1/orderBook/<span>smt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mdt_usdt">http://data.gateio.io/api2/1/orderBook/<span>mdt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mdt_eth">http://data.gateio.io/api2/1/orderBook/<span>mdt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mdt_btc">http://data.gateio.io/api2/1/orderBook/<span>mdt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mana_usdt">http://data.gateio.io/api2/1/orderBook/<span>mana_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mana_eth">http://data.gateio.io/api2/1/orderBook/<span>mana_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lun_usdt">http://data.gateio.io/api2/1/orderBook/<span>lun_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lun_eth">http://data.gateio.io/api2/1/orderBook/<span>lun_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/salt_usdt">http://data.gateio.io/api2/1/orderBook/<span>salt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/salt_eth">http://data.gateio.io/api2/1/orderBook/<span>salt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/fuel_usdt">http://data.gateio.io/api2/1/orderBook/<span>fuel_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/fuel_eth">http://data.gateio.io/api2/1/orderBook/<span>fuel_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/elf_usdt">http://data.gateio.io/api2/1/orderBook/<span>elf_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/elf_eth">http://data.gateio.io/api2/1/orderBook/<span>elf_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/drgn_usdt">http://data.gateio.io/api2/1/orderBook/<span>drgn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/drgn_eth">http://data.gateio.io/api2/1/orderBook/<span>drgn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gtc_usdt">http://data.gateio.io/api2/1/orderBook/<span>gtc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gtc_eth">http://data.gateio.io/api2/1/orderBook/<span>gtc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gtc_btc">http://data.gateio.io/api2/1/orderBook/<span>gtc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qlc_usdt">http://data.gateio.io/api2/1/orderBook/<span>qlc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qlc_btc">http://data.gateio.io/api2/1/orderBook/<span>qlc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qlc_eth">http://data.gateio.io/api2/1/orderBook/<span>qlc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dbc_usdt">http://data.gateio.io/api2/1/orderBook/<span>dbc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dbc_btc">http://data.gateio.io/api2/1/orderBook/<span>dbc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dbc_eth">http://data.gateio.io/api2/1/orderBook/<span>dbc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bnty_usdt">http://data.gateio.io/api2/1/orderBook/<span>bnty_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bnty_eth">http://data.gateio.io/api2/1/orderBook/<span>bnty_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lend_usdt">http://data.gateio.io/api2/1/orderBook/<span>lend_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lend_eth">http://data.gateio.io/api2/1/orderBook/<span>lend_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/icx_usdt">http://data.gateio.io/api2/1/orderBook/<span>icx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/icx_eth">http://data.gateio.io/api2/1/orderBook/<span>icx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btf_usdt">http://data.gateio.io/api2/1/orderBook/<span>btf_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btf_btc">http://data.gateio.io/api2/1/orderBook/<span>btf_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ada_usdt">http://data.gateio.io/api2/1/orderBook/<span>ada_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ada_btc">http://data.gateio.io/api2/1/orderBook/<span>ada_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lsk_usdt">http://data.gateio.io/api2/1/orderBook/<span>lsk_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lsk_btc">http://data.gateio.io/api2/1/orderBook/<span>lsk_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/waves_usdt">http://data.gateio.io/api2/1/orderBook/<span>waves_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/waves_btc">http://data.gateio.io/api2/1/orderBook/<span>waves_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bifi_usdt">http://data.gateio.io/api2/1/orderBook/<span>bifi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bifi_btc">http://data.gateio.io/api2/1/orderBook/<span>bifi_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mds_eth">http://data.gateio.io/api2/1/orderBook/<span>mds_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/mds_usdt">http://data.gateio.io/api2/1/orderBook/<span>mds_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dgd_usdt">http://data.gateio.io/api2/1/orderBook/<span>dgd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dgd_eth">http://data.gateio.io/api2/1/orderBook/<span>dgd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qash_usdt">http://data.gateio.io/api2/1/orderBook/<span>qash_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qash_eth">http://data.gateio.io/api2/1/orderBook/<span>qash_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qash_btc">http://data.gateio.io/api2/1/orderBook/<span>qash_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/powr_usdt">http://data.gateio.io/api2/1/orderBook/<span>powr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/powr_eth">http://data.gateio.io/api2/1/orderBook/<span>powr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/powr_btc">http://data.gateio.io/api2/1/orderBook/<span>powr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/fil_usdt">http://data.gateio.io/api2/1/orderBook/<span>fil_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcd_usdt">http://data.gateio.io/api2/1/orderBook/<span>bcd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcd_btc">http://data.gateio.io/api2/1/orderBook/<span>bcd_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/sbtc_usdt">http://data.gateio.io/api2/1/orderBook/<span>sbtc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/sbtc_btc">http://data.gateio.io/api2/1/orderBook/<span>sbtc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/god_usdt">http://data.gateio.io/api2/1/orderBook/<span>god_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/god_btc">http://data.gateio.io/api2/1/orderBook/<span>god_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcx_usdt">http://data.gateio.io/api2/1/orderBook/<span>bcx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcx_btc">http://data.gateio.io/api2/1/orderBook/<span>bcx_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/hsr_usdt">http://data.gateio.io/api2/1/orderBook/<span>hsr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/hsr_btc">http://data.gateio.io/api2/1/orderBook/<span>hsr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/hsr_eth">http://data.gateio.io/api2/1/orderBook/<span>hsr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qsp_usdt">http://data.gateio.io/api2/1/orderBook/<span>qsp_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qsp_eth">http://data.gateio.io/api2/1/orderBook/<span>qsp_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ink_btc">http://data.gateio.io/api2/1/orderBook/<span>ink_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ink_usdt">http://data.gateio.io/api2/1/orderBook/<span>ink_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ink_eth">http://data.gateio.io/api2/1/orderBook/<span>ink_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ink_qtum">http://data.gateio.io/api2/1/orderBook/<span>ink_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/med_qtum">http://data.gateio.io/api2/1/orderBook/<span>med_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/med_eth">http://data.gateio.io/api2/1/orderBook/<span>med_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/med_usdt">http://data.gateio.io/api2/1/orderBook/<span>med_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bot_qtum">http://data.gateio.io/api2/1/orderBook/<span>bot_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bot_usdt">http://data.gateio.io/api2/1/orderBook/<span>bot_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bot_eth">http://data.gateio.io/api2/1/orderBook/<span>bot_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qbt_qtum">http://data.gateio.io/api2/1/orderBook/<span>qbt_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qbt_eth">http://data.gateio.io/api2/1/orderBook/<span>qbt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qbt_usdt">http://data.gateio.io/api2/1/orderBook/<span>qbt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tsl_qtum">http://data.gateio.io/api2/1/orderBook/<span>tsl_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tsl_usdt">http://data.gateio.io/api2/1/orderBook/<span>tsl_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gnx_usdt">http://data.gateio.io/api2/1/orderBook/<span>gnx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gnx_eth">http://data.gateio.io/api2/1/orderBook/<span>gnx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/neo_usdt">http://data.gateio.io/api2/1/orderBook/<span>neo_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gas_usdt">http://data.gateio.io/api2/1/orderBook/<span>gas_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/neo_btc">http://data.gateio.io/api2/1/orderBook/<span>neo_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/gas_btc">http://data.gateio.io/api2/1/orderBook/<span>gas_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/iota_usdt">http://data.gateio.io/api2/1/orderBook/<span>iota_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/iota_btc">http://data.gateio.io/api2/1/orderBook/<span>iota_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/nas_usdt">http://data.gateio.io/api2/1/orderBook/<span>nas_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/nas_eth">http://data.gateio.io/api2/1/orderBook/<span>nas_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/nas_btc">http://data.gateio.io/api2/1/orderBook/<span>nas_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/eth_btc">http://data.gateio.io/api2/1/orderBook/<span>eth_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/etc_btc">http://data.gateio.io/api2/1/orderBook/<span>etc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/etc_eth">http://data.gateio.io/api2/1/orderBook/<span>etc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zec_btc">http://data.gateio.io/api2/1/orderBook/<span>zec_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dash_btc">http://data.gateio.io/api2/1/orderBook/<span>dash_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ltc_btc">http://data.gateio.io/api2/1/orderBook/<span>ltc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bch_btc">http://data.gateio.io/api2/1/orderBook/<span>bch_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btg_btc">http://data.gateio.io/api2/1/orderBook/<span>btg_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qtum_btc">http://data.gateio.io/api2/1/orderBook/<span>qtum_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/qtum_eth">http://data.gateio.io/api2/1/orderBook/<span>qtum_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xrp_btc">http://data.gateio.io/api2/1/orderBook/<span>xrp_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/doge_btc">http://data.gateio.io/api2/1/orderBook/<span>doge_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xmr_btc">http://data.gateio.io/api2/1/orderBook/<span>xmr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zrx_btc">http://data.gateio.io/api2/1/orderBook/<span>zrx_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/zrx_eth">http://data.gateio.io/api2/1/orderBook/<span>zrx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dnt_eth">http://data.gateio.io/api2/1/orderBook/<span>dnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dpy_eth">http://data.gateio.io/api2/1/orderBook/<span>dpy_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/oax_eth">http://data.gateio.io/api2/1/orderBook/<span>oax_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rep_eth">http://data.gateio.io/api2/1/orderBook/<span>rep_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lrc_eth">http://data.gateio.io/api2/1/orderBook/<span>lrc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lrc_btc">http://data.gateio.io/api2/1/orderBook/<span>lrc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/pst_eth">http://data.gateio.io/api2/1/orderBook/<span>pst_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcdn_eth">http://data.gateio.io/api2/1/orderBook/<span>bcdn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bcdn_usdt">http://data.gateio.io/api2/1/orderBook/<span>bcdn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tnt_eth">http://data.gateio.io/api2/1/orderBook/<span>tnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/snt_eth">http://data.gateio.io/api2/1/orderBook/<span>snt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/snt_btc">http://data.gateio.io/api2/1/orderBook/<span>snt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btm_eth">http://data.gateio.io/api2/1/orderBook/<span>btm_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/btm_btc">http://data.gateio.io/api2/1/orderBook/<span>btm_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/llt_eth">http://data.gateio.io/api2/1/orderBook/<span>llt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/snet_eth">http://data.gateio.io/api2/1/orderBook/<span>snet_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/snet_usdt">http://data.gateio.io/api2/1/orderBook/<span>snet_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/llt_snet">http://data.gateio.io/api2/1/orderBook/<span>llt_snet</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/omg_eth">http://data.gateio.io/api2/1/orderBook/<span>omg_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/omg_btc">http://data.gateio.io/api2/1/orderBook/<span>omg_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/pay_eth">http://data.gateio.io/api2/1/orderBook/<span>pay_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/pay_btc">http://data.gateio.io/api2/1/orderBook/<span>pay_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bat_eth">http://data.gateio.io/api2/1/orderBook/<span>bat_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bat_btc">http://data.gateio.io/api2/1/orderBook/<span>bat_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cvc_eth">http://data.gateio.io/api2/1/orderBook/<span>cvc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/storj_eth">http://data.gateio.io/api2/1/orderBook/<span>storj_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/storj_btc">http://data.gateio.io/api2/1/orderBook/<span>storj_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/eos_eth">http://data.gateio.io/api2/1/orderBook/<span>eos_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/eos_btc">http://data.gateio.io/api2/1/orderBook/<span>eos_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bts_usdt">http://data.gateio.io/api2/1/orderBook/<span>bts_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bts_btc">http://data.gateio.io/api2/1/orderBook/<span>bts_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tips_eth">http://data.gateio.io/api2/1/orderBook/<span>tips_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xmc_usdt">http://data.gateio.io/api2/1/orderBook/<span>xmc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/xmc_btc">http://data.gateio.io/api2/1/orderBook/<span>xmc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cs_eth">http://data.gateio.io/api2/1/orderBook/<span>cs_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/cs_usdt">http://data.gateio.io/api2/1/orderBook/<span>cs_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/man_eth">http://data.gateio.io/api2/1/orderBook/<span>man_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/man_usdt">http://data.gateio.io/api2/1/orderBook/<span>man_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/rem_eth">http://data.gateio.io/api2/1/orderBook/<span>rem_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lym_eth">http://data.gateio.io/api2/1/orderBook/<span>lym_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lym_usdt">http://data.gateio.io/api2/1/orderBook/<span>lym_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/instar_eth">http://data.gateio.io/api2/1/orderBook/<span>instar_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ont_eth">http://data.gateio.io/api2/1/orderBook/<span>ont_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ont_usdt">http://data.gateio.io/api2/1/orderBook/<span>ont_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bft_eth">http://data.gateio.io/api2/1/orderBook/<span>bft_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/bft_usdt">http://data.gateio.io/api2/1/orderBook/<span>bft_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/iht_eth">http://data.gateio.io/api2/1/orderBook/<span>iht_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/iht_usdt">http://data.gateio.io/api2/1/orderBook/<span>iht_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/senc_eth">http://data.gateio.io/api2/1/orderBook/<span>senc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/senc_usdt">http://data.gateio.io/api2/1/orderBook/<span>senc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tomo_eth">http://data.gateio.io/api2/1/orderBook/<span>tomo_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/elec_eth">http://data.gateio.io/api2/1/orderBook/<span>elec_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/ship_eth">http://data.gateio.io/api2/1/orderBook/<span>ship_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/tfd_eth">http://data.gateio.io/api2/1/orderBook/<span>tfd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/hav_eth">http://data.gateio.io/api2/1/orderBook/<span>hav_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/hur_eth">http://data.gateio.io/api2/1/orderBook/<span>hur_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lst_eth">http://data.gateio.io/api2/1/orderBook/<span>lst_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/lino_eth">http://data.gateio.io/api2/1/orderBook/<span>lino_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/swh_eth">http://data.gateio.io/api2/1/orderBook/<span>swh_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/swh_usdt">http://data.gateio.io/api2/1/orderBook/<span>swh_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dock_usdt">http://data.gateio.io/api2/1/orderBook/<span>dock_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/orderBook/dock_eth">http://data.gateio.io/api2/1/orderBook/<span>dock_eth</span></a></li>
                                </ol>
            </div>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/orderBook/eth_btc
    <span style="color:#0088cc"># Response</span>
      {
        "result": "true",
        "asks": [
                [29500,    4.07172355],
                [29499,    0.00203397],
                [29495,    1],
                [29488,    0.0672],
                [29475,    0.001]
            ],
        "bids": [
                [28001, 0.0477],
                [28000, 0.35714018],
                [28000, 2.56222976],
                [27800, 0.0015],
                [27777, 0.1]
            ]
        }
            </pre>
                <h3>返回值说明</h3>
                <pre>    asks :卖方深度
    bids :买方深度</pre>
            <br>
                </div>
                <div id="history" class="item">
                    <h4>8. 历史成交记录 API <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回最新80条历史成交记录： </p>
                <p> URL: http://data.gateio.io/api2/1/tradeHistory/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]</font>
                </p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/tradeHistory/eth_btc
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "data": [
            {
                "tradeID": "27734287",
                "date": "2017-09-29 11:52:05",
                "timestamp": "1506657125",
                "type": "buy",
                "rate": 0.1,
                "amount": 0.01,
                "total": 0.001
            }
        ],
        "elapsed": "6.901ms"
    }
            </pre>
                <p> 返回从[TID]往后的最多1000历史成交记录： </p>
                <p> URL: http://data.gateio.io/api2/1/tradeHistory/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]/[TID]</font>
                </p>
                <p> 请替换 [CURR_A] and [CURR_B] 为您需要查看的币种.</p>
                <p> 支持的兑换类型: </p>
                <ol class="type-ol">
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>btc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bch_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bch_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/eth_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>eth_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/etc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>etc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qtum_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>qtum_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ltc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ltc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dash_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dash_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zec_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>zec_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btm_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>btm_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/eos_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>eos_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/req_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>req_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/snt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>snt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/omg_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>omg_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/pay_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>pay_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cvc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>cvc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zrx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>zrx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tnt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>tnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xmr_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>xmr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xrp_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>xrp_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/doge_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>doge_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bat_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bat_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/pst_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>pst_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btg_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>btg_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dpy_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dpy_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lrc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>lrc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/storj_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>storj_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rdn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>rdn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/stx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>stx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/knc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>knc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/link_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>link_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cdt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>cdt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ae_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ae_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ae_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ae_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ae_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ae_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cdt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>cdt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rdn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rdn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/stx_eth">http://data.gateio.io/api2/1/tradeHistory/<span>stx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/knc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>knc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/link_eth">http://data.gateio.io/api2/1/tradeHistory/<span>link_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/req_eth">http://data.gateio.io/api2/1/tradeHistory/<span>req_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rcn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rcn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/trx_eth">http://data.gateio.io/api2/1/tradeHistory/<span>trx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/arn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>arn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/kick_eth">http://data.gateio.io/api2/1/tradeHistory/<span>kick_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bnt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ven_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ven_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mco_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mco_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/fun_eth">http://data.gateio.io/api2/1/tradeHistory/<span>fun_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/data_eth">http://data.gateio.io/api2/1/tradeHistory/<span>data_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rlc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rlc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zsc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>zsc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/wings_eth">http://data.gateio.io/api2/1/tradeHistory/<span>wings_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ctr_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ctr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mda_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mda_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rcn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>rcn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/trx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>trx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/kick_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>kick_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ven_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ven_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mco_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mco_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/fun_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>fun_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/data_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>data_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zsc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>zsc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mda_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mda_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xtz_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>xtz_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gnt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gnt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>gnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gem_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gem_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gem_eth">http://data.gateio.io/api2/1/tradeHistory/<span>gem_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rfr_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>rfr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rfr_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rfr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dadi_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dadi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dadi_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dadi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/abt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>abt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/abt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>abt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ledu_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ledu_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ledu_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ledu_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ledu_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ledu_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ost_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ost_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ost_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ost_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xlm_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>xlm_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xlm_eth">http://data.gateio.io/api2/1/tradeHistory/<span>xlm_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xlm_btc">http://data.gateio.io/api2/1/tradeHistory/<span>xlm_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mobi_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mobi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mobi_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mobi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mobi_btc">http://data.gateio.io/api2/1/tradeHistory/<span>mobi_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ocn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ocn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ocn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ocn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ocn_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ocn_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zpt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>zpt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zpt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>zpt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zpt_btc">http://data.gateio.io/api2/1/tradeHistory/<span>zpt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cofi_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>cofi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cofi_eth">http://data.gateio.io/api2/1/tradeHistory/<span>cofi_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/jnt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>jnt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/jnt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>jnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/jnt_btc">http://data.gateio.io/api2/1/tradeHistory/<span>jnt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/blz_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>blz_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/blz_eth">http://data.gateio.io/api2/1/tradeHistory/<span>blz_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gxs_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gxs_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gxs_btc">http://data.gateio.io/api2/1/tradeHistory/<span>gxs_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mtn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mtn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mtn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mtn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ruff_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ruff_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ruff_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ruff_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ruff_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ruff_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tnc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>tnc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tnc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tnc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tnc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>tnc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zil_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>zil_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zil_eth">http://data.gateio.io/api2/1/tradeHistory/<span>zil_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tio_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>tio_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tio_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tio_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bto_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bto_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bto_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bto_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/theta_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>theta_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/theta_eth">http://data.gateio.io/api2/1/tradeHistory/<span>theta_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ddd_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ddd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ddd_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ddd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ddd_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ddd_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mkr_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mkr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mkr_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mkr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dai_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dai_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/smt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>smt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/smt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>smt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mdt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mdt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mdt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mdt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mdt_btc">http://data.gateio.io/api2/1/tradeHistory/<span>mdt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mana_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mana_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mana_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mana_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lun_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>lun_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lun_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lun_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/salt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>salt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/salt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>salt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/fuel_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>fuel_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/fuel_eth">http://data.gateio.io/api2/1/tradeHistory/<span>fuel_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/elf_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>elf_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/elf_eth">http://data.gateio.io/api2/1/tradeHistory/<span>elf_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/drgn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>drgn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/drgn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>drgn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gtc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gtc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gtc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>gtc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gtc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>gtc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qlc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>qlc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qlc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>qlc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qlc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>qlc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dbc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dbc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dbc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>dbc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dbc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dbc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bnty_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bnty_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bnty_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bnty_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lend_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>lend_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lend_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lend_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/icx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>icx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/icx_eth">http://data.gateio.io/api2/1/tradeHistory/<span>icx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btf_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>btf_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btf_btc">http://data.gateio.io/api2/1/tradeHistory/<span>btf_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ada_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ada_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ada_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ada_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lsk_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>lsk_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lsk_btc">http://data.gateio.io/api2/1/tradeHistory/<span>lsk_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/waves_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>waves_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/waves_btc">http://data.gateio.io/api2/1/tradeHistory/<span>waves_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bifi_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bifi_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bifi_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bifi_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mds_eth">http://data.gateio.io/api2/1/tradeHistory/<span>mds_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/mds_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>mds_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dgd_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dgd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dgd_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dgd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qash_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>qash_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qash_eth">http://data.gateio.io/api2/1/tradeHistory/<span>qash_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qash_btc">http://data.gateio.io/api2/1/tradeHistory/<span>qash_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/powr_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>powr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/powr_eth">http://data.gateio.io/api2/1/tradeHistory/<span>powr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/powr_btc">http://data.gateio.io/api2/1/tradeHistory/<span>powr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/fil_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>fil_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcd_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bcd_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcd_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bcd_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/sbtc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>sbtc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/sbtc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>sbtc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/god_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>god_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/god_btc">http://data.gateio.io/api2/1/tradeHistory/<span>god_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bcx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcx_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bcx_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/hsr_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>hsr_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/hsr_btc">http://data.gateio.io/api2/1/tradeHistory/<span>hsr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/hsr_eth">http://data.gateio.io/api2/1/tradeHistory/<span>hsr_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qsp_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>qsp_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qsp_eth">http://data.gateio.io/api2/1/tradeHistory/<span>qsp_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ink_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ink_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ink_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ink_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ink_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ink_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ink_qtum">http://data.gateio.io/api2/1/tradeHistory/<span>ink_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/med_qtum">http://data.gateio.io/api2/1/tradeHistory/<span>med_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/med_eth">http://data.gateio.io/api2/1/tradeHistory/<span>med_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/med_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>med_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bot_qtum">http://data.gateio.io/api2/1/tradeHistory/<span>bot_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bot_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bot_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bot_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bot_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qbt_qtum">http://data.gateio.io/api2/1/tradeHistory/<span>qbt_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qbt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>qbt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qbt_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>qbt_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tsl_qtum">http://data.gateio.io/api2/1/tradeHistory/<span>tsl_qtum</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tsl_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>tsl_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gnx_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gnx_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gnx_eth">http://data.gateio.io/api2/1/tradeHistory/<span>gnx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/neo_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>neo_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gas_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>gas_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/neo_btc">http://data.gateio.io/api2/1/tradeHistory/<span>neo_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/gas_btc">http://data.gateio.io/api2/1/tradeHistory/<span>gas_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/iota_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>iota_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/iota_btc">http://data.gateio.io/api2/1/tradeHistory/<span>iota_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/nas_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>nas_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/nas_eth">http://data.gateio.io/api2/1/tradeHistory/<span>nas_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/nas_btc">http://data.gateio.io/api2/1/tradeHistory/<span>nas_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/eth_btc">http://data.gateio.io/api2/1/tradeHistory/<span>eth_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/etc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>etc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/etc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>etc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zec_btc">http://data.gateio.io/api2/1/tradeHistory/<span>zec_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dash_btc">http://data.gateio.io/api2/1/tradeHistory/<span>dash_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ltc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>ltc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bch_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bch_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btg_btc">http://data.gateio.io/api2/1/tradeHistory/<span>btg_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qtum_btc">http://data.gateio.io/api2/1/tradeHistory/<span>qtum_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/qtum_eth">http://data.gateio.io/api2/1/tradeHistory/<span>qtum_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xrp_btc">http://data.gateio.io/api2/1/tradeHistory/<span>xrp_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/doge_btc">http://data.gateio.io/api2/1/tradeHistory/<span>doge_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xmr_btc">http://data.gateio.io/api2/1/tradeHistory/<span>xmr_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zrx_btc">http://data.gateio.io/api2/1/tradeHistory/<span>zrx_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/zrx_eth">http://data.gateio.io/api2/1/tradeHistory/<span>zrx_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dnt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dpy_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dpy_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/oax_eth">http://data.gateio.io/api2/1/tradeHistory/<span>oax_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rep_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rep_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lrc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lrc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lrc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>lrc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/pst_eth">http://data.gateio.io/api2/1/tradeHistory/<span>pst_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcdn_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bcdn_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bcdn_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bcdn_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tnt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tnt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/snt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>snt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/snt_btc">http://data.gateio.io/api2/1/tradeHistory/<span>snt_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btm_eth">http://data.gateio.io/api2/1/tradeHistory/<span>btm_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/btm_btc">http://data.gateio.io/api2/1/tradeHistory/<span>btm_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/llt_eth">http://data.gateio.io/api2/1/tradeHistory/<span>llt_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/snet_eth">http://data.gateio.io/api2/1/tradeHistory/<span>snet_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/snet_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>snet_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/llt_snet">http://data.gateio.io/api2/1/tradeHistory/<span>llt_snet</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/omg_eth">http://data.gateio.io/api2/1/tradeHistory/<span>omg_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/omg_btc">http://data.gateio.io/api2/1/tradeHistory/<span>omg_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/pay_eth">http://data.gateio.io/api2/1/tradeHistory/<span>pay_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/pay_btc">http://data.gateio.io/api2/1/tradeHistory/<span>pay_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bat_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bat_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bat_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bat_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cvc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>cvc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/storj_eth">http://data.gateio.io/api2/1/tradeHistory/<span>storj_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/storj_btc">http://data.gateio.io/api2/1/tradeHistory/<span>storj_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/eos_eth">http://data.gateio.io/api2/1/tradeHistory/<span>eos_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/eos_btc">http://data.gateio.io/api2/1/tradeHistory/<span>eos_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bts_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bts_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bts_btc">http://data.gateio.io/api2/1/tradeHistory/<span>bts_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tips_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tips_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xmc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>xmc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/xmc_btc">http://data.gateio.io/api2/1/tradeHistory/<span>xmc_btc</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cs_eth">http://data.gateio.io/api2/1/tradeHistory/<span>cs_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/cs_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>cs_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/man_eth">http://data.gateio.io/api2/1/tradeHistory/<span>man_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/man_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>man_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/rem_eth">http://data.gateio.io/api2/1/tradeHistory/<span>rem_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lym_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lym_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lym_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>lym_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/instar_eth">http://data.gateio.io/api2/1/tradeHistory/<span>instar_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ont_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ont_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ont_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>ont_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bft_eth">http://data.gateio.io/api2/1/tradeHistory/<span>bft_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/bft_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>bft_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/iht_eth">http://data.gateio.io/api2/1/tradeHistory/<span>iht_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/iht_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>iht_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/senc_eth">http://data.gateio.io/api2/1/tradeHistory/<span>senc_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/senc_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>senc_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tomo_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tomo_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/elec_eth">http://data.gateio.io/api2/1/tradeHistory/<span>elec_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/ship_eth">http://data.gateio.io/api2/1/tradeHistory/<span>ship_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/tfd_eth">http://data.gateio.io/api2/1/tradeHistory/<span>tfd_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/hav_eth">http://data.gateio.io/api2/1/tradeHistory/<span>hav_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/hur_eth">http://data.gateio.io/api2/1/tradeHistory/<span>hur_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lst_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lst_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/lino_eth">http://data.gateio.io/api2/1/tradeHistory/<span>lino_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/swh_eth">http://data.gateio.io/api2/1/tradeHistory/<span>swh_eth</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/swh_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>swh_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dock_usdt">http://data.gateio.io/api2/1/tradeHistory/<span>dock_usdt</span></a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/tradeHistory/dock_eth">http://data.gateio.io/api2/1/tradeHistory/<span>dock_eth</span></a></li>
                                </ol>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/tradeHistory/eth_btc/3275876
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "data": [
                {
                "tradeID": 3175762,
                "date": "2017-08-25 07:24:28",
                "type": "sell",
                "rate": 29011,
                "amount": 0.0019,
                "total": 55.1209
                },
                {
                "tradeID": 3175771,
                "date": "2017-08-25 07:24:41",
                "type": "sell",
                "rate": 29011,
                "amount": 0.0032,
                "total": 92.8352
                }
            ],
        "elapsed": "31.518ms"
    }
            </pre>
                <h3>返回值说明</h3>
                <pre>    amount: 成交币种数量
    date: 订单时间
    rate: 币种单价
    total: 订单总额
    tradeID: tradeID
    type: 买卖类型, buy买 sell卖</pre>
            </div>
            <br>
                </div>
				<div id="kline" class="item">
                    <h4>9. 交易市场K线数据 <i class="caret"></i></h4>
                </div>
                <div>
            <div class="sectioncont">
                <p> 返回市场最近时间段内的K先数据： </p>
                <p> URL: http://data.gateio.io/api2/1/candlestick2/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]</font>?group_sec=<font color="red">[GROUP_SEC]</font>&amp;range_hour=<font color="red">[RANGE_HOUR]</font>
                </p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    http://data.gateio.io/api2/1/candlestick2/btc_usdt?group_sec=60&amp;range_hour=1
    <span style="color:#0088cc"># Response</span>
    {
		"result": "true",
		"data": [
			[1524556800000, 12.8756391686, 9266.02, 9266.21, 9231.01, 9233.99],
			[1524558600000, 18.3153691562, 9255.79, 9284, 9251.01, 9252],
			[1524560400000, 17.5693326481, 9255, 9290.37, 9238.89, 9255.79],
			[1524562200000, 13.0314933662, 9281.92, 9295, 9239, 9255.4],
			[1524564000000, 26.0352948098, 9290.92, 9327.46, 9282, 9282],
			[1524565800000, 24.5813921079, 9300, 9340, 9290, 9297.5],
			[1524567600000, 27.8742949471, 9345, 9348.99, 9291.91, 9300],
			[1524569400000, 29.3344591996, 9299.99, 9349.99, 9290, 9343.93],
			[1524571200000, 28.4720072519, 9305, 9327.99, 9286.32, 9290.02],
			...
		]
	}
            </pre>
                <p> 返回[range_hour]范围内的K线数据： </p>
                <p> URL: http://data.gateio.io/api2/1/candlestick2/<font color="red">[CURR_A]</font>_<font color="red">[CURR_B]?group_sec=[GROUP_SEC]&amp;range_hour=[RANGE_HOUR]</font>
                </p>
                <p> 请替换 [CURR_A] 和 [CURR_B] 为您需要查看的币种，GROUP_SEC 和 RANGE_HOUR 替换为需要获取的时间区域.</p>
                <p> 支持的兑换类型: </p>
                <ol class="type-ol">
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bch_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bch_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/eth_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>eth_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/etc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>etc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qtum_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qtum_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ltc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ltc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dash_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dash_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zec_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zec_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btm_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btm_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/eos_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>eos_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/req_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>req_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/snt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>snt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/omg_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>omg_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/pay_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>pay_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cvc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cvc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zrx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zrx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tnt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tnt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xmr_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xmr_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xrp_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xrp_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/doge_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>doge_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bat_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bat_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/pst_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>pst_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btg_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btg_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dpy_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dpy_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lrc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lrc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/storj_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>storj_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rdn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rdn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/stx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>stx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/knc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>knc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/link_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>link_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cdt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cdt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ae_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ae_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ae_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ae_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ae_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ae_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cdt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cdt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rdn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rdn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/stx_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>stx_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/knc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>knc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/link_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>link_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/req_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>req_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rcn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rcn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/trx_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>trx_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/arn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>arn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/kick_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>kick_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bnt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bnt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ven_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ven_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mco_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mco_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/fun_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>fun_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/data_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>data_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rlc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rlc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zsc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zsc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/wings_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>wings_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ctr_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ctr_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mda_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mda_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rcn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rcn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/trx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>trx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/kick_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>kick_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ven_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ven_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mco_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mco_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/fun_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>fun_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/data_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>data_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zsc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zsc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mda_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mda_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xtz_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xtz_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gnt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gnt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gnt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gnt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gem_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gem_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gem_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gem_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rfr_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rfr_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rfr_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rfr_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dadi_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dadi_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dadi_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dadi_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/abt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>abt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/abt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>abt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ledu_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ledu_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ledu_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ledu_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ledu_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ledu_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ost_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ost_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ost_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ost_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xlm_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xlm_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xlm_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xlm_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xlm_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xlm_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mobi_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mobi_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mobi_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mobi_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mobi_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mobi_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ocn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ocn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ocn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ocn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ocn_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ocn_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zpt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zpt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zpt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zpt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zpt_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zpt_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cofi_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cofi_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cofi_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cofi_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/jnt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>jnt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/jnt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>jnt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/jnt_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>jnt_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/blz_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>blz_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/blz_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>blz_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gxs_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gxs_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gxs_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gxs_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mtn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mtn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mtn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mtn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ruff_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ruff_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ruff_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ruff_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ruff_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ruff_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tnc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tnc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tnc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tnc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tnc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tnc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zil_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zil_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zil_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zil_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tio_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tio_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tio_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tio_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bto_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bto_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bto_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bto_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/theta_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>theta_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/theta_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>theta_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ddd_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ddd_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ddd_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ddd_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ddd_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ddd_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mkr_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mkr_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mkr_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mkr_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dai_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dai_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/smt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>smt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/smt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>smt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mdt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mdt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mdt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mdt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mdt_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mdt_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mana_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mana_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mana_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mana_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lun_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lun_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lun_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lun_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/salt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>salt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/salt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>salt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/fuel_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>fuel_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/fuel_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>fuel_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/elf_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>elf_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/elf_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>elf_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/drgn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>drgn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/drgn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>drgn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gtc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gtc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gtc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gtc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gtc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gtc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qlc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qlc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qlc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qlc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qlc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qlc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dbc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dbc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dbc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dbc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dbc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dbc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bnty_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bnty_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bnty_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bnty_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lend_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lend_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lend_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lend_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/icx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>icx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/icx_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>icx_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btf_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btf_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btf_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btf_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ada_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ada_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ada_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ada_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lsk_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lsk_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lsk_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lsk_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/waves_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>waves_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/waves_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>waves_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bifi_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bifi_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bifi_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bifi_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mds_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mds_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/mds_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>mds_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dgd_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dgd_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dgd_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dgd_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qash_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qash_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qash_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qash_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qash_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qash_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/powr_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>powr_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/powr_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>powr_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/powr_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>powr_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/fil_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>fil_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcd_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcd_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcd_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcd_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/sbtc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>sbtc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/sbtc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>sbtc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/god_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>god_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/god_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>god_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcx_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcx_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/hsr_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>hsr_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/hsr_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>hsr_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/hsr_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>hsr_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qsp_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qsp_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qsp_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qsp_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ink_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ink_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ink_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ink_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ink_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ink_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ink_qtum?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ink_qtum</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/med_qtum?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>med_qtum</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/med_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>med_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/med_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>med_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bot_qtum?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bot_qtum</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bot_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bot_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bot_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bot_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qbt_qtum?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qbt_qtum</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qbt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qbt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qbt_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qbt_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tsl_qtum?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tsl_qtum</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tsl_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tsl_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gnx_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gnx_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gnx_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gnx_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/neo_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>neo_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gas_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gas_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/neo_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>neo_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/gas_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>gas_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/iota_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>iota_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/iota_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>iota_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/nas_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>nas_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/nas_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>nas_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/nas_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>nas_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/eth_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>eth_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/etc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>etc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/etc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>etc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zec_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zec_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dash_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dash_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ltc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ltc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bch_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bch_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btg_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btg_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qtum_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qtum_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/qtum_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>qtum_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xrp_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xrp_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/doge_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>doge_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xmr_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xmr_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zrx_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zrx_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/zrx_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>zrx_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dnt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dnt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dpy_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dpy_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/oax_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>oax_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rep_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rep_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lrc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lrc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lrc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lrc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/pst_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>pst_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcdn_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcdn_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bcdn_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bcdn_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tnt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tnt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/snt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>snt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/snt_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>snt_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btm_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btm_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/btm_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>btm_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/llt_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>llt_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/snet_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>snet_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/snet_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>snet_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/llt_snet?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>llt_snet</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/omg_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>omg_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/omg_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>omg_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/pay_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>pay_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/pay_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>pay_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bat_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bat_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bat_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bat_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cvc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cvc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/storj_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>storj_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/storj_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>storj_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/eos_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>eos_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/eos_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>eos_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bts_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bts_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bts_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bts_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tips_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tips_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xmc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xmc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/xmc_btc?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>xmc_btc</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cs_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cs_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/cs_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>cs_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/man_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>man_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/man_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>man_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/rem_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>rem_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lym_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lym_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lym_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lym_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/instar_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>instar_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ont_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ont_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ont_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ont_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bft_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bft_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/bft_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>bft_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/iht_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>iht_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/iht_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>iht_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/senc_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>senc_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/senc_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>senc_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tomo_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tomo_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/elec_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>elec_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/ship_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>ship_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/tfd_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>tfd_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/hav_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>hav_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/hur_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>hur_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lst_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lst_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/lino_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>lino_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/swh_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>swh_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/swh_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>swh_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dock_usdt?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dock_usdt</span>?group_set=60&amp;range_hour=1</a></li>
                                    <li><a class="api-type" target="_blank" href="http://data.gateio.io/api2/1/candlestick2/dock_eth?group_sec=60&amp;range_hour=1">http://data.gateio.io/api2/1/candlestick2/<span>dock_eth</span>?group_set=60&amp;range_hour=1</a></li>
                                </ol>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    GET:&nbsp;http://data.gateio.io/api2/1/candlestick2/btc_usdt?group_sec=60&amp;range_hour=1
    <span style="color:#0088cc"># Response</span>
    {
		"result": "true",
		"data": [
			[1524556800000, 12.8756391686, 9266.02, 9266.21, 9231.01, 9233.99],
			[1524558600000, 18.3153691562, 9255.79, 9284, 9251.01, 9252],
			[1524560400000, 17.5693326481, 9255, 9290.37, 9238.89, 9255.79],
			[1524562200000, 13.0314933662, 9281.92, 9295, 9239, 9255.4],
			[1524564000000, 26.0352948098, 9290.92, 9327.46, 9282, 9282],
			[1524565800000, 24.5813921079, 9300, 9340, 9290, 9297.5],
			[1524567600000, 27.8742949471, 9345, 9348.99, 9291.91, 9300],
			[1524569400000, 29.3344591996, 9299.99, 9349.99, 9290, 9343.93],
			[1524571200000, 28.4720072519, 9305, 9327.99, 9286.32, 9290.02],
			...
		]
	}
            </pre>
                <h3>返回值说明</h3>
                <pre>    time: 时间戳
    volume: 交易量
    close: 收盘价
    high: 最高价
    low: 最低价
    open: 开盘价</pre>
            </div>
            <br>
                </div>
                <div id="trade" class="item">
                    <h4>10. 自动交易 API <i class="caret"></i></h4>
                </div>
            <div class="sectioncont">
                <p> 通过以下API，用户可以使用程序控制自动进行账号资金查询，下单交易，取消挂单。
                    <br>
                    请注意：请在您的程序中设置的HTTP请求头参数 Content-Type 为 application/x-www-form-urlencoded
                </p>
                <p> 用户首先要通过<a href="https://gate.io/myaccount/apikeys">这个链接</a>获取API接口身份认证用到的Key和Secret。
                    然后在程序中用Secret作为密码，通过SHA512加密方式签名需要POST给服务器的数据得到Sign，并在HTTPS请求的Header部分传回Key和Sign。请参考以下接口说明和例子程序进行设置。
                </p>
                <p class="item"><b>获取帐号资金余额API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/balances</p>
                <p> 参数数据提交方式：POST </p>
                <p> 参数：无 </p>
                <p> 返回数据格式：JSON </p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/balances
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "available": {
            "BTC": "1000",
            "ETH": "968.8",
            "ETC": "0",
            },
        "locked": {
            "ETH": "1"
            }
    }
            </pre>
                <h3>返回值说明</h3>
                <pre>    available: 可用各币种资金余额
    locked : 冻结币种金额</pre>
            </div>


                <p class="item"><b>获取充值地址API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/depositAddress</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table style="text-align: center">
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr style="font-size: 13px;">
                        <td>currency</td>
                        <td>String</td>
                        <td>是</td>
                        <td>币种 如(BTC, LTC)</td>
                    </tr>
                </tbody></table>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/depositAddress
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "addr": "LPXtk1kWHioP62SzfqwKbYE3Z7Wt2ujYEc",
        "message": "Sucess",
        "code": 0
    }
            </pre>
                <h3>返回值说明</h3>
                <pre>    addr: 钱包地址</pre>
                <br>
                </div>

                <p class="item"><b>获取充值提现历史API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/depositsWithdrawals</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr style="font-size: 13px;">
                        <td>start</td>
                        <td>String</td>
                        <td>否</td>
                        <td>起始UNIX时间(如 1469092370)</td>
                    </tr>
                    <tr style="font-size: 13px;">
                        <td>end</td>
                        <td>String</td>
                        <td>否</td>
                        <td>终止UNIX时间(如 1469713981)</td>
                    </tr>
                </tbody></table>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/depositsWithdrawals
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "deposits": [
                {
                "id": "c204730",
                "currency": "ETH",
                "address": "1111 1111 1111 1111 1111 1",
                "amount": "222.61",
                "txid": "210496",
                "timestamp": "1474962729",
                "status": "DONE"
                }

            ],
        "withdraws": [
                {
                "currency": "ETC",
                "address": "addr2",
                "amount": "600000000",
                "txid": "2104963",
                "timestamp": "1469092378",
                "status": "DONE"
                }
            ],
        "message": "Success"
    }
            </pre>
                <h3>返回值说明</h3>
                <pre>    deposits: 充值
    withdraws: 提现
    currency: 币种
    address: 充值/提现地址
    amount: 金额
    timestamp: 发起时间戳
    status: 记录状态 DONE:完成; CANCEL:取消; REQUEST:请求中 </pre>
                <br>
                </div>
                <p class="item"><b>下单交易买入API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/buy</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr style="font-size: 13px;">
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对(如ltc_btc,ltc_btc)</td>
                    </tr>

                    <tr style="font-size: 13px;">
                        <td>rate</td>
                        <td>String</td>
                        <td>是</td>
                        <td>价格</td>
                    </tr>
                    <tr style="font-size: 13px;">
                        <td>amount</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易量</td>
                    </tr>

                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/buy
    <span style="color:#0088cc"># Response</span>
    {
        "result":"true",
        "orderNumber":"123456",
		"rate":"1000",
		"leftAmount":"0",
		"filledAmount":"0.1",
		"filledRate":"800.00",
        "message":"Success"
    }</pre>
                <h3>返回值说明</h3>
                <pre>    orderNumber: 订单单号
    rate: 下单价格	
    leftAmount: 剩余数量
    filledAmount: 成交数量	
	filledRate: 成交价格
				</pre>
                <p>注：返回的orderNumber可用于查询，取消订单。</p>
                <br>
                </div>
                <p class="item"><b>下单交易卖出API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/sell</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对(如ltc_btc,ltc_btc)</td>
                    </tr>

                    <tr>
                        <td>rate</td>
                        <td>String</td>
                        <td>是</td>
                        <td>价格</td>
                    </tr>
                    <tr>
                        <td>amount</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易量</td>
                    </tr>

                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/sell
    <span style="color:#0088cc"># Response</span>
    {
        "result":"true",
        "orderNumber":"123456",
		"rate":"1000",
		"leftAmount":"0",
		"filledAmount":"0.1",
		"filledRate":"800.00",
        "message":"Success"
    }</pre>
                <h3>返回值说明</h3>
                <pre>    orderNumber: 订单单号
	rate: 下单价格	
    leftAmount: 剩余数量
    filledAmount: 成交数量	
	filledRate: 成交价格
				</pre>
                <p>注：返回的orderNumber可用于查询，取消订单。</p>
                <br>
                </div>
                <p class="item"><b>取消下单API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/cancelOrder</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>orderNumber</td>
                        <td>String</td>
                        <td>是</td>
                        <td>下单单号</td>
                    </tr>
                    <tr>
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对(如 ltc_btc)</td>
                    </tr>
                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/cancelOrder
    <span style="color:#0088cc"># Response</span>
   {"result":"true","message":"Success"}</pre>
                <h3>返回值说明</h3>
                <pre>    result: 是否成功 true成功 false失败
    message: 提示消息</pre>
                </div>
                <p class="item"><b>取消多个下单API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/cancelOrders</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>orders_json</td>
                        <td>String</td>
                        <td>是</td>
                        <td>下单单号和pair的json数据; 示例:[
                            {
                            "orderNumber":"7942422"
                            "currencyPair":"ltc_btc"
                            },
                            {
                            "orderNumber":"7942423"
                            "currencyPair":"ltc_btc"
                            }
                            ]
                        </td>
                    </tr>

                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/cancelOrders
    <span style="color:#0088cc"># Response</span>
   {"result":"true","message":"Success"}</pre>
                <h3>返回值说明</h3>
                <pre>    result: 是否成功 true成功 false失败
    message: 提示消息</pre>
                <br>
                </div>
                <p class="item"><b>取消所有下单API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/cancelAllOrders</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>type</td>
                        <td>int</td>
                        <td>是</td>
                        <td>下单类型(0:卖出,1:买入,-1:不限制)</td>
                    </tr>
                    <tr>
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对 (示例:ltc_btc)</td>
                    </tr>
                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/cancelAllOrders
    <span style="color:#0088cc"># Response</span>
   {"result":"true","message":"Success"}</pre>
                <h3>返回值说明</h3>
                <pre>    result: 是否成功 true成功 false失败
    message: 提示消息</pre>
                <br>
                </div>

                <p class="item"><b>获取下单状态API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/getOrder</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>orderNumber</td>
                        <td>String</td>
                        <td>是</td>
                        <td>下单单号</td>
                    </tr>
                    <tr>
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对(如:eth_btc)</td>
                    </tr>
                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/getOrder
    <span style="color:#0088cc"># Response</span>
    {
        "result":"true",
        "order":{
            "id":"15088",
            "status":"cancelled",
            "currencyPair":"eth_btc",
            "type":"sell",
            "rate":811,
            "amount":"0.39901357",
            "initialRate":811,
            "initialAmount":"1"
            },
        "message":"Success"
    }</pre>
                <h3>返回值说明</h3>
                <pre>    status: 订单状态 cancelled已取消 done已完成
    currencyPair: 交易对
    type: 买卖类型 sell卖出, buy买入
    rate: 价格
    amount: 买卖数量
    initialRate:下单价格
    initialAmount:下单量
                </pre>
                <br>
                </div>

                <p class="item"><b>获取我的当前挂单列表API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/openOrders</p>
                <p> 参数数据提交方式：POST </p>
                <p> 参数：无
                </p>
                <p> 返回数据格式：JSON </p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/openOrders
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true", "message": "Success", "code": 0,"elapsed": "6.262ms",
        "orders": [
                {
                "orderNumber": "30032151",
                "type": "buy",
                "rate": 21367.521367521,
                "amount": "0.0936",
                "total": "2000",
                "initialRate": 21367.521367521,
                "initialAmount": "0.0936",
                "filledRate": 0,
                "filledAmount": 0,
                "currencyPair": "eth_btc",
                "timestamp": "1407828913",
                "status": "open"
                }
            ]
        }</pre>
                <h3>返回值说明</h3>
                <pre>    amount: 订单总数量 剩余未成交数量
    currencyPair: 订单交易对
    filledAmount: 已成交量
    filledRate: 成交价格
    initialAmount: 下单量
    initialRate: 下单价格
    orderNumber: 订单号
    rate: 交易单价
    status: 订单状态
    timestamp: 时间戳
    total:总计
    type: 买卖类型 buy:买入;sell:卖出</pre>
                <br>
                </div>

                <p class="item"><b>获取我的24小时内成交记录API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/tradeHistory</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>

                <p> 参数：
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>currencyPair</td>
                        <td>String</td>
                        <td>是</td>
                        <td>交易币种对</td>
                    </tr>
                    <tr>
                        <td>orderNumber</td>
                        <td>String</td>
                        <td>否</td>
                        <td>订单号</td>
                    </tr>
                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/tradeHistory
    <span style="color:#0088cc"># Response</span>
    {
        "result": "true",
        "message": "Success",
        "trades": [
                {
                "id": "7942422",
                "orderid": "38100491",
                "pair": "ltc_btc",
                "type": "buy",
                "rate": "0.01719",
                "amount": "0.0588",
                "time": "06-12 02:49:11",
                "time_unix": "1402512551"
                }
            ]
    }</pre>
                <h3>返回值说明</h3>
                <pre>    orderid: 订单id
    pair: 交易对
    type: 买卖类型
    rate: 买卖价格
    amount: 订单买卖币种数量
    time: 订单时间
    time_unix: 订单unix时间戳</pre>
                <br>
                </div>
                <p class="item"><b>提现API</b></p>
                <div>
                <p> API URL: https://api.gateio.io/api2/1/private/withdraw</p>
                <p> 参数数据提交方式：POST </p>
                <p> 返回数据格式：JSON </p>
                <h3>请求参数</h3>
                <p>
                </p><table>
                    <tbody><tr style="background-color: #f5f5f5;">
                        <td width="100">参数名</td>
                        <td width="200">参数类型</td>
                        <td width="100">必填</td>
                        <td width="200">描述</td>
                    </tr>
                    <tr>
                        <td>currency</td>
                        <td>String</td>
                        <td>是</td>
                        <td>提现币种(如:btc)</td>
                    </tr>
                    <tr>
                        <td>amount</td>
                        <td>String</td>
                        <td>是</td>
                        <td>提现数量</td>
                    </tr>
                    <tr>
                        <td>address</td>
                        <td>String</td>
                        <td>是</td>
                        <td>提现地址(如:1HkxtBAMrA3tP5ENnYY2CZortjZvFDH5Cs)</td>
                    </tr>
                </tbody></table>
                <p></p>
                <h3>示例</h3>
                <pre class="prettyprint">    <span style="color:#0088cc"># Request</span>
    POST:&nbsp;https://api.gateio.io/api2/1/private/withdraw
    <span style="color:#0088cc"># Response</span>
    { "result": "true",  "message": "Success"}</pre>
                <h3>返回值说明</h3>
                <pre>    result: 是否成功 true成功 false失败
    message: 提示消息</pre>
                </div>
            </div>
                <div id="errCode" class="item">
                    <h4>11. 错误代码说明 <i class="caret"></i></h4>
                </div>
                <div>
                    <div class="sectioncont">
                        <p> 系统返回错误码对应说明 </p>
                    </div>

                    <table class="table">
                        <tbody><tr>
                            <th align="left" width="250px">错误代码</th>
                            <th align="left">详细描述</th>
                        </tr>
                        <tr>
                            <td>1</td>
                            <td>无效请求</td>
                        </tr>
                        <tr>
                            <td>2</td>
                            <td>无效版本</td>
                        </tr>
                        <tr>
                            <td>3</td>
                            <td>无效请求</td>
                        </tr>
                        <tr>
                            <td>4</td>
                            <td>请求太频繁，稍后再试</td>
                        </tr>
                        <tr>
                            <td>5,6</td>
                            <td>Key或签名无效，请重新创建</td>
                        </tr>
                        <tr>
                            <td>7</td>
                            <td>币种对不支持</td>
                        </tr>
                        <tr>
                            <td>8,9</td>
                            <td>币种不支持</td>
                        </tr>
                        <tr>
                            <td>10</td>
                            <td>验证错误</td>
                        </tr>
                        <tr>
                            <td>11</td>
                            <td>地址获取失败</td>
                        </tr>
                        <tr>
                            <td>12</td>
                            <td>参数为空</td>
                        </tr>
                        <tr>
                            <td>13</td>
                            <td>系统错误，联系管理员</td>
                        </tr>
                        <tr>
                            <td>14</td>
                            <td>无效用户</td>
                        </tr>
                        <tr>
                            <td>15</td>
                            <td>撤单太频繁，一分钟后再试</td>
                        </tr>
                        <tr>
                            <td>16</td>
                            <td>无效单号，或挂单已撤销</td>
                        </tr>
                        <tr>
                            <td>17</td>
                            <td>无效单号</td>
                        </tr>
                        <tr>
                            <td>18</td>
                            <td>无效挂单量</td>
                        </tr>
                        <tr>
                            <td>19</td>
                            <td>交易已暂停</td>
                        </tr>
                        <tr>
                            <td>20</td>
                            <td>挂单量太小</td>
                        </tr>
                        <tr>
                            <td>21</td>
                            <td>资金不足</td>
                        </tr>


                        </tbody></table>
                </div>
                <br>
				<div id="example" class="item">
                    <h4>12. 示例程序 <i class="caret"></i></h4>
                </div>
				<div>
					<div class="sectioncont">
					</div>
					<div class="sectioncont">
					   <p> PHP示例程序 (更多参考：<a style="color:#329ddc" href="https://github.com/gateio/rest/tree/master/java" target="_blank">JAVA</a>,
					   <a style="color:#329ddc" href="https://github.com/gateio/rest/tree/master/nodejs" target="_blank">NODEJS</a>,
					   <a style="color:#329ddc" href="https://github.com/gateio/rest/tree/master/python" target="_blank">PYTHON</a>,
					   <a style="color:#329ddc" href="https://github.com/gateio/rest/tree/master/go" target="_blank">GO</a>,
					   <a style="color:#329ddc" href="https://github.com/gateio/rest/tree/master/php" target="_blank">PHP</a>
					   )</p>
					</div>
					<div id="divExamplePHP" style="overflow:auto; max-height: 650px;">
					<pre class="codeSample">					<code>

	&lt;?php

		function gateio_query($path, array $req = array()) {
			// API settings, add your Key and Secret at here
			$key = '';
			$secret = '';

			// generate a nonce to avoid problems with 32bits systems
			$mt = explode(' ', microtime());
			$req['nonce'] = $mt[1].substr($mt[0], 2, 6);

			// generate the POST data string
			$post_data = http_build_query($req, '', '&amp;');
			$sign = hash_hmac('sha512', urldecode($post_data), $secret);

			// generate the extra headers
			$headers = array(
				'KEY: '.$key,
				'SIGN: '.$sign
			);

			//!!! please set Content-Type to application/x-www-form-urlencoded if it's not the default value

			// curl handle (initialize if required)
			static $ch = null;
			if (is_null($ch)) {
				$ch = curl_init();
				curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
				curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; gateio PHP bot; '.php_uname('a').'; PHP/'.phpversion().')');
			}

			curl_setopt($ch, CURLOPT_URL, 'https://api.gateio.io/api2/'.$path);
			curl_setopt($ch, CURLOPT_POSTFIELDS, $post_data);
			curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
			curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
			curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);


			// run the query
			$res = curl_exec($ch);

			if ($res === false) throw new Exception('Could not get reply: '.curl_error($ch));
			//var_dump($res);
			//print_r($res);
			$dec = json_decode($res, true);
			if (!$dec) throw new Exception('Invalid data received, please make sure connection is working and requested API exists: '.$res);

			return $dec;
		}


		function curl_file_get_contents($url) {

			// our curl handle (initialize if required)
			static $ch = null;
			if (is_null($ch)) {
				$ch = curl_init();
				curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
				curl_setopt($ch, CURLOPT_USERAGENT,
					'Mozilla/4.0 (compatible; gateio PHP bot; '.php_uname('a').'; PHP/'.phpversion().')'
					);
			}
			curl_setopt($ch, CURLOPT_URL, 'https://api.gateio.io/api2/'.$url);
			curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);

			// run the query
			$res = curl_exec($ch);
			if ($res === false) throw new Exception('Could not get reply: '.curl_error($ch));
			//echo $res;
			$dec = json_decode($res, true);
			if (!$dec) throw new Exception('Invalid data: '.$res);

			return $dec;
		}

		function get_top_rate($currency_pair, $type='BUY') {

			$url = '1/orderBook/'.strtoupper($currency_pair);
			$json = curl_file_get_contents($url);

			$rate = 0;

			if (strtoupper($type) == 'BUY') {
				$r =  $json['bids'][0];
				$rate = $r[0];
			} else  {
				$r = end($json['asks']);
				$rate = $r[0];
			}

			return $rate;
		}

		function get_pairs() {

			$url = '1/pairs';
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_marketinfo(){

			$url = '1/marketinfo';
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_tickers(){

			$url = '1/tickers';
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_ticker($current_pairs){

			$url = '1/ticker/'.strtoupper($current_pairs);
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_orderbooks(){

			$url = '1/orderBooks';
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_orderbook($current_pairs){

			$url = '1/orderBook/'.strtoupper($current_pairs);
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_trade_history($current_pairs, $tid){

			$url = '1/tradeHistory/'.strtoupper($current_pairs).'/'.$tid;
			$json = curl_file_get_contents($url);

			return $json;
		}

		function get_balances() {

			return gateio_query('1/private/balances');
		}

		function withdraw($currency, $amount, $address) {

			return gateio_query('1/private/withdraw',
				array(
					'currency' =&gt; strtoupper($currency),
					'amount' =&gt; $amount,
					'address' =&gt; $address
				)
			);
		}

		function get_order($order_number, $currency_pair) {

			return gateio_query('1/private/getOrder',
				array(
					'currencyPair' =&gt; strtoupper($currency_pair),
					'orderNumber' =&gt; $order_number
				)
			);
		}

		function cancel_order($order_number, $currency_pair) {

			return gateio_query('1/private/cancelOrder',
				array(
					'currencyPair' =&gt; strtoupper($currency_pair),
					'orderNumber' =&gt; $order_number
				)
			);
		}

		function cancel_all_orders($type, $currency_pair) {

			return gateio_query('1/private/cancelAllOrders',
				array(
					'type' =&gt; $type,
					'currencyPair' =&gt; strtoupper($currency_pair)
				)
			);
		}

		function sell($currency_pair, $rate, $amount) {

			return gateio_query('1/private/sell',
				array(
					'currencyPair' =&gt; strtoupper($currency_pair),
					'rate' =&gt; $rate,
					'amount' =&gt; $amount,
				)
			);
		}

		function buy($currency_pair, $rate, $amount) {

			return gateio_query('1/private/buy',
				array(
					'currencyPair' =&gt; strtoupper($currency_pair),
					'rate' =&gt; $rate,
					'amount' =&gt; $amount,
				)
			);
		}

		function get_my_trade_history($currency_pair, $order_number) {

			return gateio_query('1/private/tradeHistory',
				array(
					'currencyPair' =&gt; strtoupper($currency_pair),
					'orderNumber' =&gt; $order_number
				)
			);
		}

		function open_orders() {

			return gateio_query('1/private/openOrders');
		}

		function deposit_address($currency) {

			return gateio_query('1/private/depositAddress',
				array(
					'currency' =&gt; strtoupper($currency)
				)
			);
		}


		try {
			// example 1: get balances
			var_dump(get_balances());

			// example 2: place a buy order
			$pair = 'ltc_btc';
			$rate = get_top_rate($pair, $type) * 1.01;
			var_dump(buy($pair, $rate, '0.01'));

			// example 3: cancel an order
			var_dump(cancel_order(125811, $pair));

			// example 4: get order status
			var_dump(get_order(15088, $pair));

			// example 5: list all open orders
			var_dump(open_orders());

		} catch (Exception $e) {
			echo "Error:".$e-&gt;getMessage();

		}
	?&gt;
					</code>

					</pre>
					</div>
				</div>

            </div>