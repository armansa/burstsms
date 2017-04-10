<!DOCTYPE html>

<html>
<head>
  <title>Burst SMS Client</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

</head>

<body>
   {{.title}}
   <form method="post">
     <p>
       <label for="phone_number">Number</label>
       <input name="phone_number" id="phone_number" type="number"/>
     </p>
     <p>
       <label for="message_body">Message</label>
       <input name="message_body" id="message_body" type="text"/>
     </p>
     <p>
       <input type="submit">
     </p>
   </form>
</body>
</html>
