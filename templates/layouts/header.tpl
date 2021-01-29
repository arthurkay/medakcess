{{ define "header" }}
<!DOCTYPE html>
<html lang="en_us">
<head>
    <meta charset="utf-8">
    <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" name="viewport" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />

    <title>{{ .PageTitle }}</title>

    <!-- Scripts -->
    <script src="assets/js/app.js" defer></script>

    <!-- Fonts -->
    <link rel="stylesheet" href="assets/vendor/font-awesome/css/font-awesome.min.css">
    <link rel="dns-prefetch" href="https://fonts.gstatic.com">
    <link href="https://fonts.googleapis.com/css?family=Raleway:300,400,600" rel="stylesheet" type="text/css">

    <!-- Styles -->
    <link href="assets/css/app.css" rel="stylesheet">
    <!-- <link href="assets/css/sb-admin.css" rel="stylesheet"> --> 
    <link rel="stylesheet" href="assets/css/material-dashboard.css?v=2.0.0">
    <!-- Documentation extras -->
    <!-- CSS Just for demo purpose, don't include it in your project -->
    <link href="assets/css/demo.css" rel="stylesheet" />
</head>
    <div id="app">
        <nav class="navbar navbar-expand-md navbar-light navbar-laravel">
            <div class="container">
                <a class="navbar-brand" href="/">
                    Med<i class="fa fa-paper-plane-o" style="font-size:36px;color: teal;"></i>Akcess
                </a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                    <!-- Right Side Of Navbar -->
                    <ul class="navbar-nav ml-auto">
                      
                    </ul>
                
            </div>
        </nav>

 {{ end }}