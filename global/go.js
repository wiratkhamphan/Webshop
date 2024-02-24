var gloabmeun = document.querySelector("#globalmeun")
gloabmeun.innerHTML = `
        <a href="/index.html" class="lek">Home</a>
        <a href="/global/POS/pos.html">ขายสินค้า</a>
        <div class="dropdown">
            <button class="dropbtn" onclick="toggleDropdown()">Dropdown
              <i class="fa fa-caret-down"></i>
            </button>
            <div class="dropdown-content" id="myDropdown">
            <a href="/global/Stock/stock.html">Stock</a>
              <a href="#">Link 2</a>
              <a href="#">Link 3</a>
            </div>
          </div> 
        
        <a href="/global/Login/Login.html" class="logout-link">Logout</a>
`
