<?php

if ($_SERVER['REQUEST_URI'] !== '/upload') {
    http_response_code(404);
    return;
}

http_send_content_type('application/json');

if (!$_FILES['image']) {
    http_response_code(400);
    echo json_encode(array('message' => 'no file provided'));
    return;
}

$file = $_FILES['image'];

if ($file['size'] > 20 << 20) {
    http_response_code(400);
    echo json_encode(array('message' => 'file size too large'));
    return;
}

// do something with file

http_response_code(204);
