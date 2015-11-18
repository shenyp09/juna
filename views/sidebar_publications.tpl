<section class="leftNav">
    <div class="navTitle">
        {{i18n .Lang "pubs_short"}}
    </div>
    <ul class="nav nav-pills nav-stacked sidebar">
        <li><a class="{{.PublicationsActive}}" href="/pub/publications">{{i18n .Lang "pub.pub"}}</a></li>
        <li><a class="{{.ProposalsActive}}" href="/pub/proposals">{{i18n .Lang "pub.prop"}}</a></li>
        <li><a class="{{.TalksActive}}" href="/pub/talks">{{i18n .Lang "pub.talk"}}</a></li>
        <li><a class="{{.ThesesActive}}" href="/pub/theses">{{i18n .Lang "pub.thes"}}</a></li>
    </ul>
</section>
