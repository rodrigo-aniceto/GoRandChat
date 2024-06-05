
document.getElementById('form-in').onsubmit = function(e) {
    
    var checkbox = document.getElementById('agree-terms');
    if (!checkbox.checked) {
        e.preventDefault();
        alert('You must agree to the terms before logging in.');
    }
};
