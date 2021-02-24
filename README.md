# html 
html is a go library for generating html components such as buttons, form fields, forms, tables and more using golang.

This package was created from work I have done on [Beubo](https://github.com/uberswe/beubo).

A more advanced package for writing html in go is [htmlgo](https://github.com/julvo/htmlgo).

## Example

Writing something like this

```go
textField := Text().Label("Test").Class("test").Id("test").Name("test")

button := Button().Class("button").Content("This is a button")

f := Form().Class("test").Method("POST").Action("/").Fields(textField, button)

fmt.Println(f.Render())
```

Will output the following

```html
<div class="test">
    <form method="POST" action="/">
        <div class="test">
            <label for="test">Test</label>
            <input type="text" id="test" name="test" value="" placeholder="">
        </div>
        <button class="button">This is a button</button>
    </form>
</div>
```