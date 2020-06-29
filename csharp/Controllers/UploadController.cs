using System;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace csharp.Controllers
{
    [ApiController]
    [Route("/upload")]
    public class UploadController : Controller
    {
        [HttpPost]
        public IActionResult Post()
        {
            try
            {
                var file = Request.Form.Files.GetFile("image");
                if (file.Length > 20 << 20)
                {
                    return BadRequest(new
                    {
                        Message = "file too large"
                    });
                }

                return NoContent();
            }
            catch
            {
                return BadRequest(new
                {
                    Message = "missing file"
                });
            }
        }

        private void HandleFile(IFormFile file)
        {
            // do something with the file
        }
    }
}