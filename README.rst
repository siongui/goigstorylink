===================================
Web Scrape Instagram Stories in Go_
===================================

Get all links of Instagram stories in Go_.


Obtain Cookies
++++++++++++++

The following three values are must to access the Instagram story API.

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

First login to Instagram_ from Chrome browser, and there are two ways to get the
above information:

1. From `Chrome Developer Tools`_: See this `SO answer`_ or `Obtain cookies`_
   section in `instastories-backup`_ repo.
2. From Chrome extension: Use EditThisCookie_ or `cookie-txt-export`_ or other
   cookie tools.


Get User Info
+++++++++++++

See `Instagram API -Get the userId <https://stackoverflow.com/a/44773079>`_


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `Chrome IG Story — Bribing the Instagram Story API with cookies 🍪🍪🍪 <https://medium.com/@calialec/chrome-ig-story-bribing-the-instagram-story-api-with-cookies-c813e6dff911>`_

.. [2] | Python:
       | `GitHub - bl1nk/instadump: 📼 Dumps the Instagram stories of all your friends <https://github.com/bl1nk/instadump>`_
       | `GitHub - ragulbalaji/IGStoriesDumper: Instagram keeps all ur stories data far beyond when u think they do, this bulk downloads those Instagram Stories <https://github.com/ragulbalaji/IGStoriesDumper>`_
       | `GitHub - hoschiCZ/instastories-backup: Backup your friends' Instagram Stories forever and get to keep them even after 24 hours. <https://github.com/hoschiCZ/instastories-backup>`_

.. [3] | Go:
       | `GitHub - DiSiqueira/InstagramStoriesDownloader: Small software to keep watching and download all stories in your account <https://github.com/DiSiqueira/InstagramStoriesDownloader>`_
       | `GitHub - gugadev/storiesgram: Get instagram stories from third public accounts (or yours) <https://github.com/gugadev/storiesgram>`_
       | `GitHub - ahmdrz/goinsta: Unofficial Instagram API written in Golang <https://github.com/ahmdrz/goinsta>`_

.. [4] | JavaScript:
       | `GitHub - jlobos/instagram-stories: Get the Instagram Stories in Node.js and Browser <https://github.com/jlobos/instagram-stories>`_

.. [5] | Chrome Extension:
       | `GitHub - CaliAlec/ChromeIGStory: Chrome extension that lets you view your friend's Instagram Stories in the browser. <https://github.com/CaliAlec/ChromeIGStory>`_
       | `GitHub - 9paradox/InstagramDownloader: Chrome Extension  to download Instagram images and videos <https://github.com/9paradox/InstagramDownloader>`_
       | `GitHub - kryptokinght/InstagramCatch: A Chrome Extension that helps you in downloading Instagram pics and videos in one click only. <https://github.com/kryptokinght/InstagramCatch>`_
       | `GitHub - qwerty22121998/InstaDown: Download image from instagram using chrome extension <https://github.com/qwerty22121998/InstaDown>`_
       | `GitHub - davidmaillo/instagram-images-download-extension: Google Chrome extension that helps you to download all the images from an Instagram's user profile page in a click, without API. Directly from browser. A very useful tool to backup your photos or from another user. <https://github.com/davidmaillo/instagram-images-download-extension>`_
       | `GitHub - Ation/instagram_downloader: Chrome extension to download opened photo from Instagram <https://github.com/Ation/instagram_downloader>`_
       | 
       | `GitHub - cjkumaresh/instagram-downloader: download images from instagram <https://github.com/cjkumaresh/instagram-downloader>`_

.. [1] | `golang http request with cookie - Google search <https://www.google.com/search?q=golang+http+request+with+cookie>`_
       | `golang http request with cookie - DuckDuckGo search <https://duckduckgo.com/?q=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Ecosia search <https://www.ecosia.org/search?q=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Qwant search <https://www.qwant.com/?q=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Bing search <https://www.bing.com/search?q=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Yahoo search <https://search.yahoo.com/search?p=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Baidu search <https://www.baidu.com/s?wd=golang+http+request+with+cookie>`_
       | `golang http request with cookie - Yandex search <https://www.yandex.com/search/?text=golang+http+request+with+cookie>`_
       |
       | `Go HTTP Post and use Cookies - Stack Overflow <https://stackoverflow.com/questions/12756782/go-http-post-and-use-cookies>`_
       | `golang request with cookie <http://constd.com/post/golang-request-with-cookie>`_

.. _Go: https://golang.org/
.. _UNLICENSE: http://unlicense.org/
.. _Web Scrape: https://www.google.com/search?q=Web+Scrape
.. _EditThisCookie: https://www.google.com/search?q=EditThisCookie
.. _cookie-txt-export: https://github.com/siongui/cookie-txt-export.go
.. _Obtain cookies: https://github.com/hoschiCZ/instastories-backup#obtain-cookies
.. _instastories-backup: https://github.com/hoschiCZ/instastories-backup
.. _Chrome Developer Tools: https://developer.chrome.com/devtools
.. _SO answer: https://stackoverflow.com/a/44773079
