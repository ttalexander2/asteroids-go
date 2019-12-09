/**************************************************************************************************
|   Assignment:  Final Project Part 3:  shaders.go
|      Authors:  Thomas Alexander & Cameron Larson
|				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
|       Grader:  Tito Ferra & Josh Xiong
|
|       Course:  CSc 372
|   Instructor:  L. McCann
|     Due Date:  12-9-19 3:30pm
|
|  Description:  This file holds the OpenGL shader constants. These are given to the graphics context
|				 to render vertices.
|
|     Language:  GoLang
| Ex. Packages:  OpenGl, GLFW
|				 github.com/go-gl/gl/v4.1-core/gl
|				 github.com/go-gl/glfw/v3.2/glfw
|
| Deficiencies:  I know of no unsatisfied requirements and no logic errors.
**************************************************************************************************/

package main

const (
	//Shader used to render vertices
	vertexShader = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"
	//Fragment shader to render everything in white
	fragmentShaderWhite = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
	//Vertex shader used for rendering fonts
	vertexFontShader = `
		#version 330 core
		layout (location = 0) in vec4 vertex; // <vec2 pos, vec2 tex>
		out vec2 TexCoords;

		uniform mat4 projection;

		void main()
		{
			gl_Position = projection * vec4(vertex.xy, 0.0, 1.0);
			TexCoords = vertex.zw;
		}
	` + "\x00"
	//Fragment shader used for rendering fonts
	fragmentFontShader = `
	#version 330 core
	in vec2 TexCoords;
	out vec4 color;
	
	uniform sampler2D text;
	uniform vec3 textColor;
	
	void main()
	{    
		vec4 sampled = vec4(1.0, 1.0, 1.0, texture(text, TexCoords).r);
		color = vec4(textColor, 1.0) * sampled;
	} 
` + "\x00"
)
