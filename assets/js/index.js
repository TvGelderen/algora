let pageContainer = null;

window.addEventListener('load', event => {
    pageContainer = document.getElementById('page-container');

    initHeader();
});

function updatePageContent() {
    if (pageContainer) {
        pageContainer.setAttribute('hx-get', `/pages${window.location.pathname}`);
        htmx.process(pageContainer);
    }
}

function initHeader() {
    const header = document.querySelector('header');
    const menuIcon = document.querySelector('.menu-icon-wrapper');

    if (header && menuIcon) {
        const resizeListener = () => {
            if (window.innerWidth > 768) {
                header.dataset.open = "false";
                window.removeEventListener('resize', resizeListener);
            }
        }

        menuIcon.addEventListener('click', () => {
            if (header.dataset.open == "true") {
                header.dataset.open = "false";
                window.removeEventListener('resize', resizeListener);
            } else {
                header.dataset.open = "true";
                window.addEventListener('resize', resizeListener);
            }
        });
    }

    const links = header.querySelectorAll('#nav-links .nav-link');

    const updateUrl = event => {
        const newUrl = window.location.origin + event.target.dataset.href;
        window.history.pushState({ path: newUrl }, '', newUrl);

        updateLinks();
        updatePageContent();
    }

    const updateLinks = () => {
        for (const link of links) {
            if (link.innerText == "Home") {
                if (link.dataset.href == window.location.pathname) {
                    link.classList.add("active");
                } else {
                    link.classList.remove("active");
                }
            } else if (window.location.href.includes(link.dataset.href)) {
                link.classList.add("active");
            } else {
                link.classList.remove("active");
            }
        }
    }

    updateLinks();
    updatePageContent();

    for (const link of links) {
        link.addEventListener('click', updateUrl);
    }
}
