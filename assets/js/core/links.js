    function dashBoard() {
            var xr = new XMLHttpRequest();
            xr.onreadystatechange = function() {
                if (xr.readyState == 4 && xr.status == 200) {
                    document.getElementById('page-content').innerHTML  = xr.responseText;
                }
            }
            xr.open("GET", "{{route('dashboard')}}");
            xr.send();
        }

        function products() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('products')}}");
            xhr.send();
        }


        function userProfile() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('profile')}}");
            xhr.send();
        }


        function overview() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('overview')}}");
            xhr.send();
        }


        function news() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('news')}}");
            xhr.send();
        }

            function income() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('income')}}");
            xhr.send();
        }

            function expenses() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('expenses')}}");
            xhr.send();
        }

         function debtMgt() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('debtmgt')}}");
            xhr.send();
        }

         function clearDebt() {
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    document.getElementById('page-content').innerHTML = xhr.responseText;
                }
            }
            xhr.open("GET", "{{route('cleardebt')}}");
            xhr.send();
        }




        function createDebt() {
            var params = new FormData();
            var debtor = document.getElementById('debtor').value;
            var currency = document.getElementById("currency").value;
            var amount = document.getElementById('amount').value;
            var interest = document.getElementById('interest').value;
            //Adding values to the form
            params.append('debtor',debtor);
            params.append('currency', currency);
            params.append('amount', amount);
            params.append('interest', interest);

            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (xhr.status == 200 && xhr.readyState == 4) {
                    alert(xhr.responseText);
                }
            }
            xhr.open("POST", "{{route('newdebt')}}");
            xhr.setRequestHeader("X-CSRF-TOKEN", document.head.querySelector("[name=csrf-token]").content);
            xhr.send(params);
        }