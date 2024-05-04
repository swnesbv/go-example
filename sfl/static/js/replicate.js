/**
 * Add more forms v0.0.1
 * Author: Satpal Bhardwaj
 * Description: Add more forms, change each input naming and grouping
 * 
 * options: {
 *  disable-naming: boolean,
 *  wrapper: string,
 *  group: string,
 *  add-btn: string,
 *  remove-btn: string
 * }
 */

(function ($) {
  'use strict';

  const trimAttr = (attr) => attr.replace('[', '').replace(']', '');

  const setupInputNames = (wrapper, group, arrayName, disableNaming) => {
    $(wrapper).find(group).each((index, element) => {
      // add index to the group
      $(element).attr(trimAttr(group), index);

      // change the input naming to array
      $(element)
        .find('input, select')
        .each((ix, element) => {
          if (typeof $(element).attr(trimAttr(disableNaming)) === 'undefined') {
            const name = getInputName(element);
            $(element).attr('name', `${arrayName}[${index}][${name}]`);
          }
        });
    });
  };

  const getInputName = (input) => {
    let name = $(input).attr('data-name');

    if (!name) {
      name = $(input).attr('name');
      $(input).attr('data-name', name);
    }

    return name;
  };

  $.fn.replicate = function (options = {}) {
    const { disableNaming, wrapper, group, addBtn, removeBtn } = options;

    if (!wrapper || !group || !addBtn || !removeBtn) {
      throw new Error('Missing required options');
    }

    const arrayName = $(this).attr(trimAttr(wrapper));

    setupInputNames(this, group, arrayName, disableNaming);

    $(document).on('click', addBtn, function () {
      const newGroup = $(this).closest(group).clone();

      newGroup.find('input:not(:radio), select').val('');
      newGroup.find('input:radio, input:checkbox').prop('checked', false);

      $(newGroup).insertAfter($(this).closest(group));

      setupInputNames($(this).closest(wrapper), group, arrayName, disableNaming);
    });

    $(document).on('click', removeBtn, function () {
      const wrap = $(this).closest(wrapper);
      const allGroups = $(wrap).find(group);

      if (allGroups.length <= 1) {
        alert('At least one item is required!');
        return;
      }

      $(this).closest(group).remove();

      setupInputNames($(wrap), group, arrayName, disableNaming);
    });
  };
})(jQuery);
