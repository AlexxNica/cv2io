# cv2me
cv2me is a service that is currently heavily under development. It will implement the CV 2.0 specification created by the CV 2.0 - Global Resume Community Group. Its core features are available under the MIT license. If you want to know more, contact Sanja Bonic via Hangouts or e-mail.

front-end: HTML5, CSS, JavaScript

back-end: Go

via Atom text editor, console compilation

## How To Run
1. **local server**

`# cd src`

`# go run cv2.go main.go osops.go`

2. **browser**

localhost:8080

For more details, see http://localhost:8080/about and http://localhost:8080/documentation on localhost or http://cv2.me.

## How To Contribute
Please use **pull requests** if you want to contribute to this code and, if you can help it, **don't wrap your code** after a certain character limit. Just write the full line and preferably use an editor that uses automatic word wrap. Below you'll find some To Dos - contact me via Hangouts if you want to contribute more, don't know where to start, or have suggestions to improve the existing code.

If you want to contribute to this code, you have to be a member of the W3C Community Group: https://www.w3.org/community/cv2/

## Next To Dos
The To Dos are sorted by *easy*, *medium*, *hard* categories. Proper Go packages will be introduced once the code works seamlessly - for now, it will stay as is for easy Go running. Proper GitHub issues will be introduced once the existing core code issues have been improved.

- [1] *easy:* import data from LinkedIn, other networks
- [2] *easy:* prettier forms
- [3] *easy:* guide for designers on how to create templates
- [4] *easy:* proper documentation for the project
- [5] *easy:* extra process on start-up that deletes files that are on the server for more than 20 minutes - e.g. in case of server restart/crash for privacy purposes for users who are not registered
- [6] *medium:* templates - finish and fix the transfer of form data into templates
- [7] *medium:* export to .pdf and .jpg from selected templates
- [8] *medium:* template selection
- [9] *hard:* authentication module - registration, profiles
- [10] *hard:* plugin module with API
- [11] *hard:* upload module for template designers
- [12] *hard:* matching and filtering system for recruiters
