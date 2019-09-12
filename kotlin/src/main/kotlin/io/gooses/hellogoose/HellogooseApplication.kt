package io.gooses.hellogoose

import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.GetMapping

@SpringBootApplication
class HellogooseApplication

fun main(args: Array<String>) {
	runApplication<HellogooseApplication>(*args)
}

@Controller
class HelloGoose {

	@GetMapping(value = "/")
	fun index(model: Model,
			  @Value("\${cf.instance.index}") instanceIndex: String,
			  @Value("\${cf.instance.guid}") instanceGuid: String) : String {

		model.addAttribute("instance_guid", instanceGuid)
		model.addAttribute("instance_index", instanceIndex)

		return "index"
	}

}
