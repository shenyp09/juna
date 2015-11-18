<div class="content">
    <div class="container">
        <section class="section nopadding">
            <div class="row">
                <div class="col-md-3 col-sm-3 col-xs-2">
                    <div class="ydate-big">
                        <span>{{.news_day}}</span> {{.news_year}}.{{.news_month}}
                    </div>
                </div>
                <div class="col-md-9 col-sm-9 col-xs-10">
                    <section class="news-item-section">
                        <ol class="breadcrumb">
                            <li><a href="/">{{i18n .Lang "home"}}</a></li>
                            <li><a href="/news">{{i18n .Lang "news"}}</a></li>
                            <li class="active">{{.news_title}}</li>
                        </ol>
                        <div class="news-item-img">
                            <img src="{{.news_img}}" alt="">
                        </div>
                        <div class="news-item-main">
                            <div class="news-item-title">
                                {{.news_title}}
                            </div>
                            <div class="col-md-9">
                                <div id="news-item-content" class="news-item-content">
                                    {{.news_content}}
                                </div>
                            </div>
                            <div class="col-md-3">
                                <div class="news-item-info">
                                    {{.news_info}}
                                </div>
                            </div>
                        </div>
                    </section>
                </div>
            </div>
        </section>
    </div>
</div>
