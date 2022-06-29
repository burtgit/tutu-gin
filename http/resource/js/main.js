window.onscroll = function() {
    Limg();
}
window.onload = function() {
    var img = document.querySelectorAll("img[data-src]");
    for (var i = 0; i < img.length; i++) {
        img[i].style.opacity = "0";
    }
    Limg();
}

function Limg() {
    var viewHeight = document.documentElement.clientHeight;
    var t = document.documentElement.scrollTop || document.body.scrollTop;
    var limg = document.querySelectorAll("img[data-src]");
    Array.prototype.forEach.call(limg, function(item, index) {
        var rect;
        if (item.getAttribute("data-src") === "")
            return;
        rect = item.getBoundingClientRect();
        if (rect.bottom >= 0 && rect.top < viewHeight) {
            (function() {
                var img = new Image();
                img.src = item.getAttribute("data-src");
                item.src = img.src;
                var j = 0;
                setInterval(function() {
                    j += 0.2;
                    if (j <= 1) {
                        item.style.opacity = j;
                        return;
                    }
                }, 100)
                item.removeAttribute('data-src');
            })()
        }
    })
}
$(document).ready(function(){$(function(){setInterval(function(){$(".preloader").fadeOut()},500)});$({property:0}).animate({property:100},{duration:1000,step:function(){var percentage=Math.round(this.property);$('.preloader__progress span').css('width',percentage+"%");if(percentage==100){$(".preloader__progress span").addClass("done")}}})});