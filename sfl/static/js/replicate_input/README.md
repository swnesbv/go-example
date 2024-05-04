# jQuery Replicate Plugin Documentation

The jQuery Replicate plugin provides functionality to dynamically replicate and manage groups of input elements within a wrapper element. It allows you to easily add and remove groups of inputs while automatically updating the input names to support array-like naming conventions.

![final-output](Screenshot.png)

## Usage

To use the jQuery Replicate plugin, you need to include the jQuery library and the plugin file in your HTML document. For example:

```html
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
<script src="jquery.replicate.js"></script>
```

Make sure to replace `jquery.replicate.js` with the actual path to the plugin file.

## Initialization

To initialize the plugin, call the `replicate` function on a jQuery object representing the container element. For example:

```javascript
$('#container').replicate(options);
```

The `options` parameter is an object that specifies the configuration for the plugin. It must include the following properties:

- `wrapper`: A string specifying the attribute name of the wrapper element. This attribute will contain the name of the array.
- `group`: A string specifying the selector for the group of elements to be replicated.
- `addBtn`: A string specifying the selector for the "Add" button that triggers the addition of a new group.
- `removeBtn`: A string specifying the selector for the "Remove" button that triggers the removal of a group.

### Example Initialization

```javascript
$('#container').replicate({
  wrapper: 'data-wrapper',
  group: '.group',
  addBtn: '.add-button',
  removeBtn: '.remove-button'
});
```

## Input Naming

The plugin automatically updates the input names when adding or removing groups. It uses array-like naming conventions for the inputs to facilitate handling the data on the server-side. By default, the input names are set to `[arrayName][index][inputName]`.

For example, if the `wrapper` attribute is set to `data-wrapper="myArray"`, and the input name is `inputName`, the resulting input name for the first group will be `myArray[0][inputName]`.

### Disabling Naming

You can disable the automatic naming for specific inputs by adding the attribute specified in the `disableNaming` option. If an input has this attribute, its name will not be updated when adding or removing groups.

### Data Attribute Cache

To prevent loss of input names during replication, the plugin utilizes the `data-name` attribute. If an input element does not have a `data-name` attribute, the plugin caches the original `name` attribute value for later use.

## Events

The plugin binds the necessary event handlers to the document to handle the "Add" and "Remove" button clicks. When the "Add" button is clicked, a new group is cloned and inserted after the current group. When the "Remove" button is clicked, the corresponding group is removed.

## Error Handling

If any of the required options (`wrapper`, `group`, `addBtn`, `removeBtn`) is missing when initializing the plugin, an error will be thrown.

## License

This plugin is released under the MIT License. See the `LICENSE` file for more information.

## Example HTML Structure

```html
<div id="container" data-wrapper="myArray">
  <div class="group">
    <!-- Input elements -->
    <input type="text" name="inputName" />
    <!-- ... -->
    <button class="remove-button">Remove</button>
  </div>
  <button class="add-button">Add</button>
</div>
```

In the above example, the `container` element represents the wrapper element. The `group` class identifies the group of elements to be replicated. The "Add" and "Remove" buttons are specified with the `add-button` and `remove-button` classes, respectively.

**Note:** Ensure that you have unique selectors for the `addBtn` and `removeBtn` options to avoid conflicts with other elements in your HTML document.

That's it! You should now have a basic understanding of how to use the jQuery Replicate plugin. Enjoy replicating and managing groups of input elements effortlessly!

# Contributor
[Satpal Bhardwaj](https://sbsharma.com/javascript/)

# Feel free to reach me for any query
<a target="_blank" href="https://www.facebook.com/Sbsharma-2798360506847821"><img src="https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white"></a>
<a target="_blank" href="https://twitter.com/Ss101Bhardwaj"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white"></a>
<a target="_blank" href="https://www.linkedin.com/in/satpal-bhardwaj-5a76b4134"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"></a>
<a target="_blank" href="https://codepen.io/sb_sharma"><img src="https://img.shields.io/badge/Codepen-000000?style=for-the-badge&logo=codepen&logoColor=white"></a>
