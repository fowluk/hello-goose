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
		String hostname = System.getenv("INSTANCE_GUID");
		model.addAttribute("hostname", hostname);
                return "index";
        }

	public static void main(String[] args) {
		SpringApplication.run(HelloGoose.class, args);
	}
}
