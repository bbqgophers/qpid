package http

const statusT = `<html>
<body>
QPID<br/>
Time: {{.Time}} <br/>
Sensors:<br/>
{{ range $idx, $sensor := .GrillSensors }}
Location: {{ $sensor.Location }} <br/>
Description: {{ $sensor.Description }}
{{ $t, e := $sensor.Temperature }}
Temperature: {{ $t.F}} <br/>
{{ end }}
</body>
</html>`
