function toggleDropdown() {
    
        document.getElementById("myDropdown").classList.toggle("show");
      
      }
      
      window.onclick = function(e) {
        if (!e.target.matches('.dropbtn')) {
          var myDropdown = document.getElementById("myDropdown");
          if (myDropdown.classList.contains('show')) {
            myDropdown.classList.remove('show');
          }
        }
      }
      
// JavaScript
var globalmenu = document.querySelector("#globalmeun");
globalmenu.innerHTML = `
<div class="navbar">
<div class="navbar-links">
  <a href="/index.html">HOME</a>
  <a href="/global/POS/pos.html">ขายสินค้า</a>
  <div class="dropdown">
    <button class="dropbtn" onclick="toggleDropdown()">Stock
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content" id="myDropdown">
      <a href="/global/Stock/stock.html">โปรแกรมยิงกล่องรับเข้า</a>
      <a href="#">โปรแกรมยิงรับเข้า</a>
      <a href="#">รายงาน Stock</a>
    </div>
  </div>
</div>
<div class="">
  <a href="#">รายงาน</a>
</div>
<div class="logout-link">
  <a href="/global/Login/Login.html">Logout</a>
</div>
</div>


`;



