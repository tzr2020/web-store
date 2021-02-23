!
function(f, h) {
    "object" == typeof exports && "undefined" != typeof module ? module.exports = h() : "function" == typeof define && define.amd ? define(h) : f.scrollnav = h()
} (this,
function() {
    function f(f, h) {
        var y, w = {};
        for (y in f) Object.prototype.hasOwnProperty.call(f, y) && (w[y] = f[y]);
        for (y in h) Object.prototype.hasOwnProperty.call(h, y) && (w[y] = h[y]);
        return w
    }
    function h(f, h) {
        if ("object" != typeof f) return Promise.reject(new Error("First argument must be an object"));
        if ("object" != typeof(h = h || document.body)) return Promise.reject(new Error("Second argument must be an object"));
        var y = h.getBoundingClientRect();
        return f.getBoundingClientRect().top - y.top
    }
    function y(f, w, E) {
        void 0 === E && (E = "ax-scrollnav");
        var L = [];
        return E += "-",
        f.forEach(function(f, O) {
            var x = [],
            j = function(f, h) {
                if ("object" != typeof f) return Promise.reject(new Error("First argument must be an object"));
                var y = f.id;
                if (!y) {
                    if ("string" != typeof h) return Promise.reject(new Error("Second argument must be a string"));
                    f.id = y = h
                }
                return y
            } (f, E + (O + 1));



            w.subSections && f.matches(w.sections) && (x = y(function(f, h, y) {
                var w = [];
                for (f = f.nextElementSibling; f && !f.matches(h);) ! y || f.matches(y) ? (w.push(f), f = f.nextElementSibling) : f = f.nextElementSibling;
                return w
            } (f, w.sections, w.subSections), w, j));
            L.push({
                id: j,
                text: f.innerText || f.textContent,
                offsetTop: h(f),
                subSections: x
            })
        }),
        L
    }
    function w(f) {
        var h = document.createElement("div");//axui
        return h.className = "ax-scrollnav",//axui
        h.innerHTML = function f(h, y) {
            void 0 === y && (y = !1);
            var w = "ax-scrollnav" + (y ? "-sub-": "-"),
            E = "\n    " + h.map(function(h) {
                return '<li class="' + w + 'item" data-section="' + h.id + '">\n            <a class="' + w + 'link" href="#' + h.id + '">' + h.text + "</a>\n            " + (h.subSections && h.subSections.length ? "" + f(h.subSections, !0) : "") + "\n          </li>"
            }).join("") + "\n  ";
            return '\n    <ol class="' + w + 'list">\n      ' + E + "\n    </ol>\n " ;
        } (f),
        h
    }
    function E(f) {
        return f.forEach(function(f) {
            var y = document.querySelector("#" + f.id);
            f.offsetTop = h(y),
            f.subSections.length && (f.subSections = E(f.subSections))
        }),
        f
    }
    function L(f, h) {
        var y = f.getAttribute("href");
        return "#" === y.charAt(0) && (y = y.substr(1)),
        function f(h, y) {
            var w;
            h.forEach(function(h) {
                h.id === y && (w = h),
                h.subSections && void 0 === w && (w = f(h.subSections, y))
            });
            return w
        } (h, y).offsetTop
    }
    var O, x, j, _ = function(f) {
        return function(h) {
            return Math.pow(h, f)
        }
    },
    I = function(f) {
        return function(h) {
            return 1 - Math.abs(Math.pow(h - 1, f))
        }
    },
    Q = function(f) {
        return function(h) {
            return h < .5 ? _(f)(2 * h) / 2 : I(f)(2 * h - 1) / 2 + .5
        }
    },
    C = {
        linear: Q(1),
        easeInQuad: _(2),
        easeOutQuad: I(2),
        easeInOutQuad: Q(2),
        easeInCubic: _(3),
        easeOutCubic: I(3),
        easeInOutCubic: Q(3),
        easeInQuart: _(4),
        easeOutQuart: I(4),
        easeInOutQuart: Q(4),
        easeInQuint: _(5),
        easeOutQuint: I(5),
        easeInOutQuint: Q(5)
    };
    function M(f, h) {
        return new Promise(function(y, w) {
            if ("number" != typeof f) return w(new Error("First argument must be a number"));
            if ("string" != typeof(h = h || "linear")) return w(new Error("Second argument must be a string"));
            var E, L = window.pageYOffset,
            O = f - L,
            x = function(f) {
                var h = Math.abs(f / 2);
                return Math.min(Math.max(h, 250), 1200)
            } (O),
            j = 20,
            _ = 0; !
            function f() {
                E = C[h]((_ += j) / x),
                window.scroll(0, E * O + L),
                _ < x ? setTimeout(f, j) : y(window.pageYOffset)
            } ()
        })
    }
    function q(f) {
        function h() {
            var h = window.scrollY || window.pageYOffset || document.body.scrollTop,
            y = h + .4 * window.innerHeight,
            w = function f(h, y, w) {
                var E, L;
                h.forEach(function(f) {
                    f.offsetTop > w ? !E && f.offsetTop < y && (E = f) : E = f
                }),
                E && E.subSections.length && (L = f(E.subSections, y, w)) && (E = L);
                return E
            } (f.data, h, y);
            return function(f, h) {
                var y = h.querySelector("[data-active]");
                if (f) {
                    var w = h.querySelector("[data-section=" + f.id + "]");
                    w && w !== y && (y && (y.classList.remove("ax-active"), y.removeAttribute("data-active")), w.classList.add("ax-active"), w.setAttribute("data-active", !0))
                } else y && (y.classList.remove("ax-active"), y.removeAttribute("data-active"))
            } (w, f.nav),
            w
        }
        return window.addEventListener("scroll", h),
        h
    }
    function B(f) {
        return f instanceof Element
    }
    return {
        init: function(h, _) {
            if (this.settings = f({
                sections: "h2",
                subSections: "h3",//axui
                insertTarget: h,
                insertLocation: "append",
                easingStyle: "easeOutQuad",
                updateHistory: !0
            },
            _), B(h)) if (!this.settings.insertTarget || B(this.settings.insertTarget)) if (["append", "prepend", "after", "before"].includes(this.settings.insertLocation)) {
                var I, Q, C, F, R = h.querySelectorAll(this.settings.sections);
                if (R.length) return this.data = y(R, this.settings),
                this.nav = w(this.data),
                Q = (I = this).settings.insertTarget,
                "append" === (C = I.settings.insertLocation) ? Q.appendChild(I.nav) : "prepend" === C ? Q.insertBefore(I.nav, Q.firstChild) : "before" === C ? Q.parentNode.insertBefore(I.nav, Q) : "after" === C && Q.parentNode.insertBefore(I.nav, Q.nextSibling),
                O = function(f) {
                    var h = f.settings;
                    function y(y) {
                        y.preventDefault();
                        var w = .39 * window.innerHeight;
                        return M(L(y.target, f.data) - w, h.easingStyle).then(function() {
                            h.updateHistory && history.replaceState({},
                            "", y.target.getAttribute("href")),
                            h.onScroll && h.onScroll()
                        })
                    }
                    return f.nav.querySelectorAll("a").forEach(function(f) {
                        f.addEventListener("click", y)
                    }),
                    y
                } (this),
                x = q(this),
                j = function(f) {
                    function h() {
                        f.data = E(f.data)
                    }
                    return window.addEventListener("resize", h),
                    h
                } (this),
                this.settings.debug && ((F = document.createElement("div")).className = "snDebugger", F.setAttribute("style", "\n      position: fixed;\n      top: 40%;\n      height: 0px;\n      border-bottom:5px solid red;\n      border-top: 5px solid blue;\n      width: 100%;\n      opacity: .5;\n      pointer-events: none;\n    "), document.body.appendChild(F)),
                this.settings.onInit ? this.settings.onInit() : void 0;
                this.settings.debug && console.error('\n        scrollnav build failed, could not find any "' + this.settings.sections + '"\n        elements inside of "' + h + '"\n      ')
            } else this.settings.debug && console.error('\n        scrollnav build failed, options.insertLocation "' + this.settings.insertLocation + '" is not a valid option\n      ');
            else this.settings.debug && console.error('\n        scrollnav build failed, options.insertTarget "' + h + '" is not an HTML Element\n      ');
            else this.settings.debug && console.error('\n        scrollnav build failed, content argument "' + h + '" is not an HTML Element\n      ')
        },
        destroy: function(h) {
            if (this.settings = f(this.settings, h),
            function(f, h) {
                f.querySelectorAll("a").forEach(function(f) {
                    f.removeEventListener("click", h)
                })
            } (this.nav, O),
            function(f) {
                window.removeEventListener("scroll", f)
            } (x),
            function(f) {
                window.removeEventListener("resize", f)
            } (j), this.nav.remove(), this.settings.onDestroy) return this.settings.onDestroy()
        },
        updatePositions: function(h) {
            if (this.settings = f(this.settings, h), this.data = E(this.data), this.settings.onUpdatePositions) return this.settings.onUpdatePositions()
        }
    }
});
//# sourceMappingURL=scrollnav.min.umd.js.map