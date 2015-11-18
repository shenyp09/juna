$(function() {
    $('img').each(function(index, el) {
        // alert($(this).attr("src"))
        if ($(this).attr("src") == "") {
            if ($(this).parent().hasClass('news-item-img')) {
                $(this).parent().remove()
            } else {
                $(this).remove()
            }
        };
    });
    var h5 = document.getElementsByTagName('body')[0].getElementsByTagName('h5')
    var len = h5.length;
    for (var i = 0; i < len; i++) {
        var header = h5[i].getElementsByTagName('a')[0]
        $clamp(header, {
            clamp: 2,
            // useNativeClamp: false
        });
    }

    $('#focusList > li > .tn > p').each(function(index, el) {
        var p = $(this)[0]
        $clamp(p, {
            clamp: 6,
            // useNativeClamp: false
        })
        var content = $(this).text()
        $(this).text('').append(content)
        $(this).children('p').each(function(index, el) {
            var content = $(this).text()
            $(this).text('').append(content)

        })
    })
    $('#newsList > li > .tn > p').each(function(index, el) {
        var p = $(this)[0]
        $clamp(p, {
                clamp: 5,
                // useNativeClamp: false
            })
            // var content = $(this).text()
            // $(this).text('').append(content)
    })

    var content = $('#news-item-content').text()
    $('#news-item-content').text('').append(content).children('p').each(function(index, el) {
        var content = $(this).text()
        $(this).text('').append(content)
    });

    
})
