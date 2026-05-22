import type { Attributes } from "graphology-types";
import { NodeProgram, type ProgramInfo } from "sigma/rendering";
import type { NodeDisplayData, RenderParams } from "sigma/types";
import { floatColor } from "sigma/utils";

const FRAGMENT_SHADER_SOURCE = `
precision mediump float;

varying vec4 v_color;

void main(void) {
  gl_FragColor = v_color;
}
`;

const VERTEX_SHADER_SOURCE = `
attribute vec4 a_id;
attribute vec4 a_color;
attribute vec2 a_position;
attribute float a_size;

uniform float u_sizeRatio;
uniform float u_pixelRatio;
uniform mat3 u_matrix;

varying vec4 v_color;

const float bias = 255.0 / 254.0;

void main() {
  gl_Position = vec4(
    (u_matrix * vec3(a_position, 1)).xy,
    0,
    1
  );

  gl_PointSize = max(1.0, a_size / u_sizeRatio * u_pixelRatio * 2.0);

  #ifdef PICKING_MODE
  v_color = a_id;
  #else
  v_color = a_color;
  #endif

  v_color.a *= bias;
}
`;

const UNIFORMS = ["u_sizeRatio", "u_pixelRatio", "u_matrix"] as const;

export class NodeSquareProgram<
  N extends Attributes = Attributes,
  E extends Attributes = Attributes,
  G extends Attributes = Attributes,
> extends NodeProgram<(typeof UNIFORMS)[number], N, E, G> {
  getDefinition() {
    return {
      VERTICES: 1,
      VERTEX_SHADER_SOURCE,
      FRAGMENT_SHADER_SOURCE,
      METHOD: WebGLRenderingContext.POINTS,
      UNIFORMS,
      ATTRIBUTES: [
        {
          name: "a_position",
          size: 2,
          type: WebGLRenderingContext.FLOAT,
        },
        {
          name: "a_size",
          size: 1,
          type: WebGLRenderingContext.FLOAT,
        },
        {
          name: "a_color",
          size: 4,
          type: WebGLRenderingContext.UNSIGNED_BYTE,
          normalized: true,
        },
        {
          name: "a_id",
          size: 4,
          type: WebGLRenderingContext.UNSIGNED_BYTE,
          normalized: true,
        },
      ],
    };
  }

  processVisibleItem(
    nodeIndex: number,
    startIndex: number,
    data: NodeDisplayData,
  ): void {
    const array = this.array;
    array[startIndex++] = data.x;
    array[startIndex++] = data.y;
    array[startIndex++] = data.size;
    array[startIndex++] = floatColor(data.color);
    array[startIndex++] = nodeIndex;
  }

  setUniforms(
    { sizeRatio, pixelRatio, matrix }: RenderParams,
    { gl, uniformLocations }: ProgramInfo<(typeof UNIFORMS)[number]>,
  ): void {
    gl.uniform1f(uniformLocations.u_pixelRatio, pixelRatio);
    gl.uniform1f(uniformLocations.u_sizeRatio, sizeRatio);
    gl.uniformMatrix3fv(uniformLocations.u_matrix, false, matrix);
  }
}
