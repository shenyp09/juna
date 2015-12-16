<section class="events">
    <div class="container">
        <div class="section">
            <div class="col-md-6 col-sm-6">
                <div class="title">{{i18n .Lang "links"}}</div>
                <ul id="linkList">
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://www.ciae.ac.cn/">CIAE</a></li>
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://www.tsinghua.edu.cn/">Tsinghua</a></li>
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://www.sjtu.edu.cn/">SJTU</a></li>
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://www.nsfc.gov.cn/">NSFC</a></li>
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://esic.ciae.ac.cn">ESIC</a></li>
                    <li class="col-md-4 col-sm-4 col-xs-4"><a href="http://esic.ciae.ac.cn/omeg2015">OMEG2015</a></li>
                </ul>
            </div>
            <div class="col-md-6 col-sm-6">
                <div class="title"><span class="more-small disable">{{i18n .Lang "more"}}</span>{{i18n .Lang "events"}}</div>
                <ul id="eventList">
                    <li class="col-md-4 col-sm-4 col-xs-4">{{.Event1}}</li>
                    <li class="col-md-4 col-sm-4 col-xs-4">{{.Event2}}</li>
                </ul>
            </div>
        </div>
    </div>
</section>
<footer>
    <div class="container">
        <section class="copyrights">
            <span class="info">
                <span>{{i18n .Lang "tel_short"}}: +86 010-69357993</span>
            <span>{{i18n .Lang "email"}}: <a href="mailto:webmaster@esic.ac.cn">webmaster@esic.ac.cn</a></span>
            <div class="clear"></div>
            <span class="copy">
            <span>©2015 Jinping Underground lab for Nuclear Astrophysics. All Rights Reserved</span>
            <span>{{i18n .Lang "visit"}}:<span id="total">{{.Visit}}</span></span> <span>{{i18n .Lang "pageview"}}:<span id="total">{{.PageView}}</span></span>
            </span>
            </span>
        </section>
    </div>
</footer>
