<section class="leftNav">
    <div class="navTitle">
        {{i18n .Lang "about"}} JUNA
    </div>
    <ul class="nav nav-pills nav-stacked sidebar">
        <li><a class="{{.GeneralActive}}" href="/about/general">{{i18n .Lang "general"}}</a></li>
        <li><a class="{{.CollaborationActive}}" href="/about/collaboration">{{i18n .Lang "coll"}}</a></li>
        <!-- <li><a class="{{.TechActive}}" href="/about/tech">{{i18n .Lang "tech"}}</a></li> -->
        <li class="disabled"><a class="{{.AnnualActive}}" href="#">{{i18n .Lang "annu"}}</a></li>
    </ul>
</section>
