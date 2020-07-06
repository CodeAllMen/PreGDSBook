function clickTraining(backgroundBig, backgroundMedium, backgroundSmall) {
    if (clickTraining && backgroundBig && backgroundMedium && backgroundSmall) {
        var screenWidth = (window.innerWidth);
        var count = 0;

        var backgroundImages = backgroundSmall;
        var padding = 'top';
        if (screenWidth >= 800) {
            backgroundImages = backgroundBig;
            padding = '50px';
        } else if (screenWidth >= 480) {
            backgroundImages = backgroundMedium;
        }

        $('#form_click_submit').on('click', function(e) {
            if (count < backgroundBig.length) {
                // replace background
                $('#call2action_before').css('background', 'url(' + backgroundImages[count] + ')  no-repeat center ' + padding);

                count++;
                return false;
            }

            return true;
        });
    }
}

$(function() {
    campaignLoaded();
    formMoAutoSubmit();

    // scroll user down
    let autoHide = $('header').data('auto-hide');

    if (autoHide === 1) {
        let scrollDown = $('header').height();
        $('html, body').animate({
            scrollTop: scrollDown + 'px'
        }, 300);
    }

    var options = function() {
        const optionsElements = document.querySelector('#form_mccmnc_select');

        if (null !== optionsElements) {
            return optionsElements.length;
        } else {
            return null;
        }
    };

    if (1 === options()) {
        document.forms[0].submit();
    }

    if(document.getElementById('ivr_form') != undefined) {
        autoClick2Call();
    }
});

$(document).ready(function() {

    var ctaInput = document.querySelector('.js-cta-input');
    if (null !== ctaInput) {
        ctaInput.addEventListener('keyup', validateForm);
    }

    var ctaCheckbox = document.querySelector('.js-cta-checkbox');
    if (null !== ctaCheckbox) {
        ctaCheckbox.addEventListener('change', validateForm);
    }


    if ($("#prelander_overlay").length > 0) {
        $("#prelander_confirm").click(function() {
            createCookie($('html').data('campaign') + '-prelander-visited', Math.floor(Date.now() / 1000));
            createCookie('isVisited', true);
            $("#prelander_overlay").fadeOut();
        });
        $("#prelander_leave").click(function() {
            window.location = $(this).data('redirect');
        });
    }
});


function validateForm() {

    var inputElement = document.querySelector('.js-cta-input');
    var checkboxElement = document.querySelector('.js-cta-checkbox');
    var buttonElement = document.querySelector('.js-cta-button');
    var messageElement = document.querySelector('.js-cta-message');

    var minLength = inputElement.getAttribute('minlength');
    var maxLength = inputElement.getAttribute('maxlength');

    var inputValidity = inputElement.checkValidity();

    messageElement.style.display = 'none';

    if (true === inputValidity) {
        inputElement.classList.add('valid');
        inputElement.classList.remove('invalid');

        if ((null !== checkboxElement && checkboxElement.checkValidity()) || null === checkboxElement) {
            buttonElement.classList.add('active');
        } else {
            buttonElement.classList.remove('active');
        }

    } else {
        inputElement.classList.add('invalid');
        inputElement.classList.remove('valid');
        buttonElement.classList.remove('active');

        if (minLength >= inputElement.value.length && maxLength <= inputElement.value.length) {
            messageElement.style.display = 'block';
        }
    }
}

function campaignLoaded() {
    var msisdn = $('#form_msisdn_input').val();
    var countryPrefix = $("#form_msisdn_input").data('country-prefix');
    var countryNetCode = $("#form_msisdn_input").data('country-netcode');

    if (typeof countryPrefix !== 'undefined') {
        countryPrefix = countryPrefix.toString();
        countryNetCode = countryNetCode.toString();

        if (typeof msisdn !== 'undefined' && msisdn.startsWith(countryPrefix)) {
            if (countryNetCode.length > 0 && countryNetCode.slice(0, 1) == '0') {
                msisdn = '0' + msisdn.slice(countryPrefix.length);
            } else {
                msisdn = msisdn.slice(countryPrefix.length);
            }

            $('#form_msisdn_input').val(msisdn);
        }
    }
}

function autoSubmitForm(form, buttonId, interval = 1000) {
    let maxCount = 20;
    let count = 0;
    count = getCookie('session_counter');

    let formSubmitInterval = setInterval(function () {
        count++;
        console.log('Interval = ' + count);
        if (count == maxCount) {
            clearInterval(formSubmitInterval);
            document.getElementById(buttonId).style.display = 'block';
        } else {
            setCookie('session_counter', count);
            form.submit();
        }
    }, interval);
}

function formMoAutoSubmit() {
    let click2sms = document.getElementById("click2sms_link");

    if (click2sms) {
        let form = $('#form_mo');
        let totalCount = 20;

        let count = 0;
        count = getCookie('session_counter');

        if (typeof form !== 'undefined') {
            setTimeout(function() {
                if (count < totalCount) {
                    count++;
                    setCookie('session_counter', count);

                    form.submit();
                } else {
                    $('#call2action_mo_submit').show();
                    $('#click2sms_link').hide();
                }
            }, 10000);
        }
    }
}

// Create cookie
function setCookie(name, value, days) {
    var expires;
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toGMTString();
    } else {
        expires = "";
    }
    document.cookie = name + "=" + value + expires + "; path=/";
}

// Read cookie
function getCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1, c.length);
        }
        if (c.indexOf(nameEQ) === 0) {
            return c.substring(nameEQ.length, c.length);
        }
    }
    return null;
}

// Erase cookie
function eraseCookie(name) {
    createCookie(name, "", -1);
}

$("#form_msisdn_input").focus(function() {
    setTimeout(function() {
        var toElement = "#form_msisdn_submit";
        var offset = 0;

        if (getMobileOperatingSystem() === 'iOS') {
            offset = $(window).height() - 273 - 35;
        } else {
            offset = $(window).height() - $(toElement).height() - 5;
        }

        $('html, body').animate({
            scrollTop: $(toElement).offset().top - offset
        });
    }, 500);
});

function getMobileOperatingSystem() {
    var userAgent = navigator.userAgent || navigator.vendor || window.opera;

    // Windows Phone must come first because its UA also contains "Android"
    if (/windows phone/i.test(userAgent)) {
        return "Windows Phone";
    }

    if (/android/i.test(userAgent)) {
        return "Android";
    }

    // iOS detection from: http://stackoverflow.com/a/9039885/177710
    if (/iPad|iPhone|iPod/.test(userAgent) && !window.MSStream) {
        return "iOS";
    }

    return "unknown";
}

function idShow(idShow) {
    $(idShow).show();
}

function idHide(idHide) {
    $(idHide).hide();
}

function autoClick2Call() {
    if(document.getElementById('ivr_form') != undefined){
        setTimeout(() => {
            document.querySelectorAll('#ivr_form .button').forEach((e) => { e.click(); });
        }, 2000);
    }
}

function createCookie(name, value, days) {
    var expires;

    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toGMTString();
    } else {
        expires = "; expires=0";
    }
    document.cookie = encodeURIComponent(name) + "=" + encodeURIComponent(value) + expires + "; path=/";
}

function readCookie(name) {
    var nameEQ = encodeURIComponent(name) + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ')
            c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) === 0)
            return decodeURIComponent(c.substring(nameEQ.length, c.length));
    }
    return null;
}

$("#form_click").submit(function(event) {
    $("#form_click_submit").prop("disabled", true);
});

function closeCampaign(broker) {
    if(broker == 'Macrokiosk'){
        return location.replace('https://google.com');
    }
    return location.replace('/');
}

(function($) {
    $.fn.portalQuiz = function(options) {
        // settings can be chnaged by implementation
        var settings = $.extend({
            pages: [{
                title: 'Welcome to the master quiz',
                image: 'https://via.placeholder.com/500x300?text=500x300',
                answerRows: 1,
                question: 'Are you ready to start the quiz?',
                answers: ['start quiz']
            }, {
                title: 'page 2',
                image: 'https://via.placeholder.com/500x300?text=500x300',
                answerRows: 2,
                question: 'Which animal is the largest?',
                answers: ['Elephant', 'Blue whale', 'Lion', 'Black bear']
            }],
            loading: {
                title: 'Loading',
                checks: ['gathering answers', 'reviewing', 'selecting your price']
            },
            color: '#000000',
            backgroundColor: '#ffffff',
            redirect: null,
            timer: {
                use: true,
                timeInSeconds: 30,
                timeAdded: 10
            }
        }, options);

        window.progresswidth = 0;

        // stylesheet changes
        $('#portal-quiz-wrapper').css('background-color', settings.backgroundColor);
        $('#portal-quiz-wrapper').css('color', settings.color);

        // handle of the function
        $(this).prepend('<div id="portal-quiz-wrapper"></div>');

        let activeStyle = 'block';
        settings.pages.forEach(function(page, index, arr) {
            if (index !== 0) {
                activeStyle = 'none';
            }

            // base page
            let html = '<div style="display: ' + activeStyle + ';" id="page-' + index + '" class="page" data-page="' + index + '">';

            html += '<div class="header-wrapper">';

            html += '<div class="timer">' + settings.timer.timeInSeconds + 's</div>';

            html += '<div class="progress-bar-wrapper">';

            html += '<div class="progress-bar">' + '<span class="bar">' + '<span class="progress"></span>' + '</span>' + '</div>';

            html += '</div>';
            html += '</div>';

            html += '<h1>' + page.title + '</h1>' + '<img class="portal-quiz-main-image" src="' + page.image + '" />';

            // question place
            html += '<h2>' + page.question + '</h2>';

            if (page.answers.length === 1) {
                page.answerRows = 1;
            }

            // answer loop
            html += '<ul class="answers" style="column-count: ' + page.answerRows + '">';
            page.answers.forEach(function(answer) {
                html += '<li class="answer">';
                if (answer.image) {
                    html += '<img class="portal-quiz-answer-image" src="' + answer.image + '" />';
                }
                if (answer.text) {
                    html += answer.text;
                }
                html += '</li>';
            });
            html += '</ul>';

            // close page
            html += '</div>';

            // add page to wrapper
            $('#portal-quiz-wrapper').append(html);
        });

        let finalPage = '<div style="display: none;" id="page-final" class="page">' + '<h1>' + settings.loading.title + '</h1>' + '<div class="circle-loader">' + '<div class="checkmark draw"></div>' + '</div>';

        // answer loop
        finalPage += '<ul>';
        settings.loading.checks.forEach(function(check, index) {
            finalPage += '<li style="display: none;" class="check" id="check-' + index + '">' + check + '</li>';
        });
        finalPage += '</ul>';

        // close page
        finalPage += '</div>';

        // add page to wrapper
        $('#portal-quiz-wrapper').append(finalPage);

        // timer
        if (settings.timer.use === true) {
            window.timer = settings.timer.timeInSeconds;

            setInterval(function() {
                if (window.timer === 0) {
                    $('.page').hide();
                    endQuiz();
                } else {
                    window.timer--;
                }

                $('.timer').html(timer + 's');
            }, 1000);
        }

        return this.find('li.answer').click(nextPage);

        function nextPage() {
            let nextPage = ($(this).closest('.page').data('page') + 1);
            $('.page').hide();

            if (settings.pages.length == nextPage) {
                endQuiz();
            } else {
                if (settings.timer.use === true) {
                    if (settings.timer.timeAdded > 0) {
                        window.timer = window.timer + settings.timer.timeAdded;
                        $('.timer').html(timer + 's');

                        let timeAddedHtml = '<div class="timeAddedHtml">+' + settings.timer.timeAdded + 's</div>';

                        $('body').prepend(timeAddedHtml);
                        $(".timeAddedHtml").fadeTo("fast", 1, function() {
                            setTimeout(function() {
                                $(".timeAddedHtml").remove();
                            }, 500);
                        });

                    }

                    window.progresswidth = window.progresswidth + (100 / settings.pages.length);

                    $('.progress').attr('style', 'width: ' + window.progresswidth + '%');
                }
                $('#page-' + nextPage).show();
            }

            $(this).parent();
        }

        function endQuiz() {
            $('#page-final').show();
            let count = 0;
            setInterval(function() {
                $('.check').hide();
                if ((settings.pages.length + 1) > count) {
                    $('#check-' + count).show();
                    count++;
                }
            }, 500);

            setTimeout(function() {
                $('.circle-loader').toggleClass('load-complete');
                $('.checkmark').toggle();

                setTimeout(function() {
                    if (settings.redirect === null) {
                        $('#portal-quiz-wrapper').remove();
                    } else {
                        window.location.href = settings.redirect;
                    }

                }, 500);
            }, ((settings.loading.checks.length + 1) * 500));
        }

    }
    ;
}(jQuery));
