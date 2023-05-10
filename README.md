# 🌟 lens


lens is like 'ls' for humans: human readable filesizes, emojis to represent filetypes, symlinks, and directories. 


## 🚀 usage 
your best bet is to download a release that corresponds with your os and architecture, then move `lens` to `/usr/bin/lens` or in your local bin folder, then alias it to `ll` or something. it's nice, but using it as a drop in replacement for `ls` could conflict with scripts that use `ls` expecting the output of `ls`.


**view the current directory:**
```
./lens
```
**or view a specific directory**

```
./lens /path/to/directory
```

the output should look something like this:

```
╭─dh@lisa ~/projects/lens ‹master●› 
╰─$ lens
dir: /users/dh/projects/lens
.git        📁 dir        384      tue, 09 may 2023 21:14:41 cdt
.gitignore  📄 gitignore  34       tue, 09 may 2023 11:26:47 cdt
go.mod      📄 mod        48       tue, 09 may 2023 11:26:47 cdt
*lens       ⚙️ exe       2107010   tue, 09 may 2023 21:14:41 cdt
license.md  📄 md         1068     tue, 09 may 2023 11:26:47 cdt
main.go     📄 go         1712     tue, 09 may 2023 20:58:44 cdt
readme.md   📄 md         877      tue, 09 may 2023 21:08:56 cdt
```

## 📄 license 
mit license 2023 donuts-are-good. license.md for more details.