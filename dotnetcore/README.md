# Hello Goose - Dot Net Core (Razor Pages)

This is a Microsoft Dot Net Core application, but you can still happily play with this on linux or macos. Dotnetcore runs happily in all three environments. This was developed and tested locally on Ubuntu.

Like the JSP version, this uses the INSTANCE_GUID environment variable rather than gethostname().

Most of what you see is the Microsoft app template, have a look in [hello-goose/Pages](hello-goose/Pages) for the custom code. The image is in [hello-goose/wwwroot](hello-goose/wwwroot).

To use it, run `cf push`.

The initial template was created with `dotnet new razor -n hello-goose` and then removing everything from the `Pages` and `wwwroot` subdirectories.
