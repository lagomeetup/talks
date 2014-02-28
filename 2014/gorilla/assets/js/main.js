(function (global, $) {
  var $simple = $('#simple');
  var $simpleLink = $('#simple-link');
  var $complex = $('#complex');
  var $complexLink = $('#complex-link');

  function toggleDisplay($el) {
    var current = 'none';
    return preventDefault(function () {
      if (current === 'none') {
        current = 'block';
      } else {
        current = 'none';
      }
      $el.css('display', current);
    });
  }

  function preventDefault(func) {
    return function (event) {
      event.preventDefault();
      func.apply(this, arguments);
    };
  }

  $simpleLink.on('click', toggleDisplay($simple));
  $complexLink.on('click', toggleDisplay($complex));

  var $addPhone = $('#add-phone');
  var $phone = $('#phone');
  var count = 0;
  $addPhone.on('click', preventDefault(function () {
    count += 1;
    var html = '<input type="text" name="Phone.' + count + '.Number" ' +
      'class="form-control" placeholder="phone number"> <input type="text" ' +
      'name="Phone.' + count + '.Kind" class="form-control" ' +
      'placeholder="phone type">';
    $phone.append(html);
  }));
}(this, this.jQuery));
