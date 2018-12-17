using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace hello_goose.Pages
{
    public class IndexModel : PageModel
    {

        public string InstanceId = "";
        public string InstanceIndex = "";

        public void OnGet()
        {
          InstanceId = Environment.GetEnvironmentVariable("CF_INSTANCE_GUID");
          InstanceIndex = Environment.GetEnvironmentVariable("CF_INSTANCE_INDEX");
        }
    }
}
