<header>
    <div class="container">
        <div class="logo">
            <img src="/static/img/style_img/logo.svg" alt="">
            {{range .RestLangs}}
            <a class="version lang-changed" style="cursor: pointer" data-lang="{{.Lang}}" class="lang-changed">{{i18n $.Lang .Name}}</a>
            {{end}}
        </div>
        <nav>
            <ul class="nav nav-pills">
                <li><a href="/">{{i18n .Lang "home"}}</a></li>
                <li role="presentation" class="dropdown"><a href="/about" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-delay="0" data-close-others="true" data-hover-delay="50" href="#" role="button" aria-expanded="false">{{i18n .Lang "about"}}</a>
                    <ul class="dropdown-menu" role="menu">
                        <li><a href="/about/general">{{i18n .Lang "general"}}</a></li>
                        <li><a href="/about/collaboration">{{i18n .Lang "coll"}}</a></li>
                        <!-- <li class="disabled"><a href="/about/tech">{{i18n .Lang "tech"}}</a></li> -->
                        <li class="disabled"><a href="#">{{i18n .Lang "annu"}}</a></li>
                    </ul>
                </li>
                <li><a href="/news">{{i18n .Lang "news"}}</a></li>
                <li><a href="/pub">{{i18n .Lang "pubs"}}</a></li>
                <li><a href="/committee">{{i18n .Lang "comm"}}</a></li>
            </ul>
        </nav>
    </div>
</header>
