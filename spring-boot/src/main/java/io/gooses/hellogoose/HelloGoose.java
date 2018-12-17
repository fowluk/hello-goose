package io.gooses.hellogoose;

import org.springframework.boot.*;
import org.springframework.boot.autoconfigure.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;

@Controller
@EnableAutoConfiguration
public class HelloGoose {

	@GetMapping(value = "/")
        public String index(Model model) {
		String instanceGuid = System.getenv("CF_INSTANCE_GUID");
		String instanceIndex = System.getenv("CF_INSTANCE_INDEX");

		model.addAttribute("instance_guid", instanceGuid);
		model.addAttribute("instance_index", instanceIndex);
                return "index";
        }

	public static void main(String[] args) {
		SpringApplication.run(HelloGoose.class, args);
	}
}
