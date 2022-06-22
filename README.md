# Console prettification project

Every time, when you create a console application, you are potentially and quite often do interact with another human being who is a user of that tool. Even if you not planning to. To have a meaningful, convenient and intuitive console user interface is still a responsibility of the instrument developer. And to use _fmt.Println()_ as a primary tool for text UI creation is not a perfect choice. While some tools do not require any elaborate UI, as the output of that tool is in fact an input of other tool, but sometimes, console output of some program do actually intended for a human consumption. And therefore, _tconsole_ golang module will help you to create something you can call "a text UI"

## Which type of the UI, you can create with tconsole ?

Or let say, what is tconsole "is not". _tconsole_ module is not a "text User Interface" with text menu, windows and other elements of full screen tUI. This module rather help you to replace _fmt.Println()_ with something more meaningful.

![tconsole UI](./example/tconsole.png "tconsole")
