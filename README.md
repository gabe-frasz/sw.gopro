# GoPro

Simple Go projects to practice.

## TODO

### tasks

- [x] Change CSV file to SQLite
- [x] Format output using `tabwriter`
- [ ] Colorize output
- [ ] Add optional due date
- [x] Display human friendly dates with `timediff`
- [x] Document CLI commands

### calculator-api

- [x] Input validation
- [ ] Error feedback
- [ ] Logging
- [ ] CORS
- [ ] Rate limiting
- [ ] Token authentication
- [ ] Tracking of calculations (DB)
- [ ] Support for floating point numbers

### qslsp

- [ ] use [gopher-lua](https://github.com/yuin/gopher-lua) to serve a real LSP server for Lua
- [ ] `textDocument/formatting`
- [ ] `textDocument/references`
- [ ] `workspace/symbol`
- [ ] `textDocument/rename`

#### How to enable custom LSP in Neovim

```lua
local client = vim.lsp.start_client({
    name = "qƨlsp",
    cmd = { "$HOME/path/to/local/qƨlsp" },
})

if not client then
       vim.notify("qƨlsp failed to start")
       return
end

vim.api.nvim_create_autocmd("FileType", {
       pattern = "markdown",
       callback = function()
               vim.lsp.buf_attach_client(0, client)
       end,
})
```

## Aknowledgements

- [Dreams of Code YouTube video](https://youtu.be/gXmznGEW9vo?si=p1nQa3W_12A3vuEI)
- [LSP from scratch with TJ](https://youtu.be/YsdlcQoHqPY?si=jADPDp8WSPzkAdyE)
