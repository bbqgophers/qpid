package http

const statusT = `<html>
<body>
QPID<br/>
Time: {{.Time}} <br/>
Sensors:<br/>
{{ range $idx, $sensor := .GrillSensors }}
&nbsp;Location: {{ $sensor.Location.String }} <br/>
&nbsp;Description: {{ $sensor.Description }} <br/>
&nbsp;Temperature: {{ $sensor.Temperature }} <br/>
{{ end }}
</body>
</html>`
