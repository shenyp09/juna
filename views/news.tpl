<div class="content">
    <div class="grey-banner">
        <div class="container banner_top">
            <img src="/static/img/style_img/banner_about.jpg" alt="">
        </div>
    </div>
    <div class="container">
        <section class="">
            <div class="col-md-3">
                <section class="leftNav">
                    <div class="navTitle">
                        {{i18n .Lang "news"}}
                    </div>
                </section>
            </div>
            <div class="col-md-9">
                <section class="newsList-all">
                    <ol class="breadcrumb">
                        <li><a href="/">{{i18n .Lang "home"}}</a></li>
                        <li class="active">{{i18n .Lang "news"}}</li>
                    </ol>
                    <ul id="newsList-all">
                        {{.Content | news}}
                    </ul>
                </section>
            </div>
        </section>
    </div>
</div>
