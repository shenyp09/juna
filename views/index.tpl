<article class="content">
    <div class="grey-banner">
        <div class="container">
            <div id="carousel-example-generic" class="carousel slide" data-ride="carousel" data-interval="8000">
                <!-- Indicators -->
                <ol class="carousel-indicators">
                    <li data-target="#carousel-example-generic" data-slide-to="0" class="active"></li>
                    <li data-target="#carousel-example-generic" data-slide-to="1"></li>
                    <li data-target="#carousel-example-generic" data-slide-to="2"></li>
                </ol>
                <!-- Wrapper for slides -->
                <div class="carousel-inner" role="listbox">
                    <div class="item active">
                        <a href="{{.banner1_link}}">
                            <img src="{{.banner1_img}}" alt="">
                            <div class="carousel-caption">
                                {{.banner1_caption}}
                            </div>
                        </a>
                    </div>
                    <div class="item">
                        <a href="{{.banner2_link}}">
                            <img src="{{.banner2_img}}" alt="">
                            <div class="carousel-caption">
                                {{.banner2_caption}}
                            </div>
                        </a>
                    </div>
                    <div class="item">
                        <a href="{{.banner3_link}}">
                            <img src="{{.banner3_img}}" alt="">
                            <div class="carousel-caption">
                                {{.banner3_caption}}
                            </div>
                        </a>
                    </div>
                </div>
                <!-- Controls -->
                <a class="left carousel-control" href="#carousel-example-generic" role="button" data-slide="prev">
                    <span class="arrow-left" aria-hidden="true"></span>
                    <span class="sr-only">Previous</span>
                </a>
                <a class="right carousel-control" href="#carousel-example-generic" role="button" data-slide="next">
                    <span class="arrow-right" aria-hidden="true"></span>
                    <span class="sr-only">Next</span>
                </a>
            </div>
        </div>
    </div>
    <div class="container">
        <section class="mainWrap">
            <section class="news section">
                <span class="title">
                <a class="more" href="news/">{{i18n .Lang "more"}}</a>
                {{i18n .Lang "news"}}
                </span>
                <ul id="newsList" class="newsList clearfix" id="newjunaen_news">
                    <li class="col-md-4 col-sm-6">
                        <div class="ydate">
                            <span>{{.news1_day}}</span> {{.news1_year}}.{{.news1_month}}
                        </div>
                        <div class="tn">
                            <h5><a href="{{.news1_link}}">{{.news1_title}}</a></h5>
                            <p id="news1_summary">{{.news1_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="ydate">
                            <span>{{.news2_day}}</span> {{.news2_year}}.{{.news2_month}}
                        </div>
                        <div class="tn">
                            <h5><a href="{{.news2_link}}">{{.news2_title}}</a></h5>
                            <p id="news2_summary">{{.news2_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="ydate">
                            <span>{{.news3_day}}</span> {{.news3_year}}.{{.news3_month}}
                        </div>
                        <div class="tn">
                            <h5><a href="{{.news3_link}}">{{.news3_title}}</a></h5>
                            <p id="news3_summary">{{.news3_summary}}</p>
                        </div>
                    </li>
                </ul>
            </section>
            <section class="focus section">
                <span class="title">{{i18n .Lang "focus"}}</span>
                <ul id="focusList" class="focusList clearfix">
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus1_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus1_link}}">{{.focus1_title}}</a></h5>
                            <p id="focus1_summary">{{.focus1_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus2_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus2_link}}">{{.focus2_title}}</a></h5>
                            <p id="focus2_summary">{{.focus2_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus3_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus3_link}}">{{.focus3_title}}</a></h5>
                            <p id="focus3_summary">{{.focus3_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus4_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus4_link}}">{{.focus4_title}}</a></h5>
                            <p id="focus4_summary">{{.focus4_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus5_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus5_link}}">{{.focus5_title}}</a></h5>
                            <p id="focus5_summary">{{.focus5_summary}}</p>
                        </div>
                    </li>
                    <li class="col-md-4 col-sm-6">
                        <div class="tn">
                            <div class="focusList-img">
                                <img src="{{.focus6_img}}" alt="">
                            </div>
                            <h5><a href="{{.focus6_link}}">{{.focus6_title}}</a></h5>
                            <p id="focus6_summary">{{.focus6_summary}}</p>
                        </div>
                    </li>
                </ul>
            </section>
        </section>
    </div>
</article>
